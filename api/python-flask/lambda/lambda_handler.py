import awsgi
from flask import (
    Flask,
    jsonify,
)

app = Flask(__name__)


@app.route('/')
def index():
    return 'Simple Flask Benchmark'

def handler(event, context):
    return awsgi.response(app, event, context)