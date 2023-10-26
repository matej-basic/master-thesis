from flask import Flask
import os

app = Flask(__name__)
app_port = os.getenv('PORT')

@app.route('/benchmark')
def helloworld():
    return "Simple Python Flask benchmark"

if __name__ == "__main__":
    app.run(port=app_port)