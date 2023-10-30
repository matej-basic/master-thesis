package main

import (
	"os"
	"os/exec"
	"log"
	"fmt"
	"net/http"
	"encoding/csv"
	"strings"
	"strconv"
	"time"
	"bufio"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Service struct {
	SERVICE_NAME	string
	SERVICE_URL		string
}

type TestingResult struct {
	TIMESTAMP			int64
	SERVICE				Service
	TIME_NAMELOOKUP		float64
	TIME_CONNECT		float64
	TIME_APPCONNECT		float64
	TIME_PRETRANSFER	float64
	TIME_REDIRECT		float64
	TIME_STARTTRANSFER	float64
	TIME_TOTAL			float64
}

var collection *mongo.Collection
var ctx 							= context.TODO()
var mongoURI 						= os.Getenv("MONGO_CONNECTION_URI") 
var accessKey 						= os.Getenv("ACCESS_KEY") 
var secretAccessKey 				= os.Getenv("SECRET_ACCESS_KEY") 
var deployment 						= os.Getenv("DEPLOYMENT")
var id 								= os.Getenv("ID")

func init() {

	// Run ID must be set to identify different deployments
	_, present := os.LookupEnv("ID")
	if (!present) {
		os.Exit(3)
	}

	// Do not initialize anything if testing
	// as I don't save testing to DB
	if deployment == "" {
		deployment = "TESTING"
		return
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("master-thesis-db").Collection("testing-results")

	
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readCsvFile(filePath string) []Service {
	var services []Service
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV file at path ", filePath)
	}

	for i, line := range data {
		if i > 0 {
			var service Service
			for j, field := range line {
				if j == 0 {
					service.SERVICE_NAME = field
				} else if j == 1 {
					service.SERVICE_URL = field
				}
			}
			services = append(services, service)
		}
	}
	return services
}

func checkLinkStatus(service Service) bool {
	fmt.Println("Checking status of service: " + service.SERVICE_NAME)
	response, err := http.Get(service.SERVICE_URL)
	if err != nil {
		log.Println("Failed to get URL: ", service.SERVICE_URL)
		return false
	}
	defer response.Body.Close()

	statusCode := response.StatusCode

	fmt.Println("Status code received from service ", service.SERVICE_NAME, " : ", statusCode)

	if statusCode == 200 {
		return true
	} else {
		return false
	}
}

func curlTesting(service Service) TestingResult {
	out, err := exec.Command("curl", "-w" ,"@curl-format.txt", "-o", "/dev/null", "-s", service.SERVICE_URL).Output()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Convert stdout from bytes[] to string
	strResult := string(out)

	// Split string into array of strings with newline character as delimiter
	temp := strings.Split(strResult, "\n")

	// Parse each index in the array into new result variable
	var result TestingResult
	result.TIMESTAMP				= time.Now().Unix()
	result.SERVICE					= service
	result.TIME_NAMELOOKUP, _ 		= strconv.ParseFloat(temp[0], 32)
	result.TIME_CONNECT, _ 			= strconv.ParseFloat(temp[1], 32)
	result.TIME_APPCONNECT, _ 		= strconv.ParseFloat(temp[2], 32)
	result.TIME_PRETRANSFER, _ 		= strconv.ParseFloat(temp[3], 32)
	result.TIME_REDIRECT, _ 		= strconv.ParseFloat(temp[4], 32)
	result.TIME_STARTTRANSFER, _ 	= strconv.ParseFloat(temp[5], 32)
	result.TIME_TOTAL, _ 			= strconv.ParseFloat(temp[6], 32)

	return result
}

func saveResultToFile(res TestingResult, startStamp string) {
	file, err := os.OpenFile("./testing_results_" + startStamp + "_" + id + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = fmt.Fprintf(writer, "%d,%s,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f\n", res.TIMESTAMP, res.SERVICE.SERVICE_NAME, res.TIME_NAMELOOKUP, res.TIME_CONNECT, res.TIME_APPCONNECT, res.TIME_PRETRANSFER, res.TIME_REDIRECT, res.TIME_STARTTRANSFER, res.TIME_TOTAL)
	check(err)
	writer.Flush()
}

func saveResultToDB(res TestingResult) error{
	_, err := collection.InsertOne(ctx, res)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func uploadResultsToS3(startStamp string) {
	s3Config := &aws.Config{
		Region: aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(accessKey,secretAccessKey, ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	filename := "./testing_results_" + startStamp + "_" + id + ".txt"
	file, _ := os.Open(filename)
	input := &s3manager.UploadInput{
		Bucket: aws.String("master-thesis-logs"),
		Key: aws.String(filename),
		Body: file,
		ContentType: aws.String("text/plain"),
	}

	output, err := uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}

func main() {
	links := readCsvFile("./service_links.txt")
	startStamp := time.Now().Format("2006-Jan-02T15-04-05")
	for _, link := range links {
		fmt.Printf("Testing %s\n", link)
		res := curlTesting(link)
		saveResultToFile(res, startStamp)

		if deployment == "PRODUCTION" {
			saveResultToDB(res) 
		}
	}
	uploadResultsToS3(startStamp)
}
