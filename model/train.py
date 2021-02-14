import codecs
import json

import pandas as pd
import xgboost
from sklearn.metrics import mean_squared_error
from sklearn.model_selection import train_test_split
from feature_engine.encoding import OrdinalEncoder
import numpy as np

if __name__ == '__main__':
    # load data
    df = pd.read_csv("test_data.csv")
    df.head()

    # Regularize data set
    df.price_per_size = df.price_per_size / 10000
    df.price = df.price / 1000000
    df.rent = df.rent / 1000

    # Test train split
    X = df.drop(columns=['price'], axis=1)
    Y = df['price']
    X_train, X_test, y_train, y_test = train_test_split(X, Y, test_size=0.33)

    # Encoding the regions
    regions_df = np.asarray(X['region']).reshape(1, -1)
    enc = OrdinalEncoder(encoding_method='ordered', variables=['region'])
    enc.fit(X_train, y_train)
    X_train_enc = enc.transform(X_train)
    X_test_enc = enc.transform(X_test)

    # fit model no training data
    regressor = xgboost.XGBRegressor(
        n_estimators=100,
        reg_lambda=1,
        gamma=0,
        max_depth=3
    )
    regressor.fit(X_train_enc, y_train)

    # make predictions for test data
    y_pred = regressor.predict(X_test_enc)

    predictions = [round(value) for value in y_pred]
    # evaluate predictions
    mse = mean_squared_error(y_test, predictions)
    print("Mean Square Error: %.2f%%" % mse)

    # Save model and encoder
    np.save('regions.npy', enc.encoder_dict_)

    with open('regions.json', 'w') as fp:
        json.dump(enc.encoder_dict_, fp)

    regressor.save_model('hemnet-pred.model')
