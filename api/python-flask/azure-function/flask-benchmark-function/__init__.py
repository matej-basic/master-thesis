import logging
from flask import Flask
import azure.functions as func

app = Flask(__name__)

def main(req: func.HttpRequest, context: func.Context) -> func.HttpResponse:
    logging.info('Python HTTP trigger function processed a request.')
    return func.WsgiMiddleware(app.wsgi_app).handle(req, context)

@app.route('/benchmark')
def helloworld():
    return "Simple Python Flask benchmark"