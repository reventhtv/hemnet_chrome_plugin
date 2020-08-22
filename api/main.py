import flask
from flask import jsonify, request
from flask_cors import CORS
import xgboost as xgb
from xgboost import Booster

app = flask.Flask(__name__)
CORS(app)
xgb_model = xgb.XGBClassifier(objective="reg:squarederror", random_state=42)

booster = Booster({'nthread': 4})
booster.load_model('poc.model')
xgb_model._Booster = booster


def make_response(payload):
    return jsonify({"results": payload})


@app.route('/predict', methods=['POST'])
def predict():
    query = request.json['query']
    # TODO: 1. Load the model using pickle
    #  2. Add more fields to request body
    #  3.
    return make_response(query)


@app.route("/health", methods=["GET"])
def health():
    # TODO: Check if can predict fasttext before sending 200
    response = flask.Response()
    response.headers["Access-Control-Allow-Origin"] = "*"
    return response


app.run(host="0.0.0.0", port=5000, debug=True)
