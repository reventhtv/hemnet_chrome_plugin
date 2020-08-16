import flask
from flask import jsonify, request
from flask_cors import CORS

app = flask.Flask(__name__)
CORS(app)


def make_response(payload):
    return jsonify({"results": payload})


@app.route('/predict', methods=['POST'])
def predict():
    query = request.json['query']
    return make_response("")


@app.route("/health", methods=["GET"])
def health():
    # TODO: Check if can predict fasttext before sending 200
    response = flask.Response()
    response.headers["Access-Control-Allow-Origin"] = "*"
    return response


app.run(host="0.0.0.0", port=5000, debug=True)
