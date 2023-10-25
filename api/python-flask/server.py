from flask import Flask

app = Flask(__name__)

@app.route('/benchmark')
def helloworld():
    return "Simple Python Flask benchmark"

application = app