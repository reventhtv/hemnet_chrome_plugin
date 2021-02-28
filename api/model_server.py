import flask
import xgboost
from flask import jsonify, request, Response
from pandas import np

app = flask.Flask(__name__)
app.config["DEBUG"] = True
bst = xgboost.XGBRegressor()


@app.route('/predict', methods=['POST'])
def server():
    re = request.json
    if "queries" in re:
        re.get("queries")
        # "queries" should be two dimensional array
        requested_query = np.array(re.get("queries"))
        # TODO: Make it thread safe
        response = bst.predict(requested_query)
        return jsonify({"result": response.tolist()})
    else:
        return Response("{'err':'bad request body'}", status=400, mimetype='application/json')


if __name__ == '__main__':
    bst.load_model('model/house.model')
    app.run()
