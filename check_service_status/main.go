package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/csv"
)

type Service struct {
	SERVICE_NAME	string
	SERVICE_URL		string
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
		log.Fatal("Failed to get URL: ", service.SERVICE_URL)
	}
	//defer response.Body.Close()

	statusCode := response.StatusCode
	data, err := ioutil.ReadAll(response.Body)

	fmt.Println("Status code received from service ", service.SERVICE_NAME, " : ", statusCode)

	if statusCode == 200 {
		fmt.Printf("Responde body received from %s : %s\n", service.SERVICE_NAME, data)
		return true
	} else {
		return false
	}
}

func main() {
	links := readCsvFile("./service_links.txt")
	for _, link := range links {
		checkLinkStatus(link)
	}
}