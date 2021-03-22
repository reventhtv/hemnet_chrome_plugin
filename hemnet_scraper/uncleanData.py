import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression

linreg = LinearRegression()
df = pd.read_csv("result-not-clean.csv")
df_unclean = pd.read_csv("result-not-clean.csv")
print(df_unclean.head(10))


sum_unclean = df_unclean.isnull().sum()
#print(sum_unclean)

df_yearlyrent_sum = df_unclean['YearlyRent'].count()
#print(df_yearlyrent_sum)

data_with_null = df_unclean[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent']].dropna()

rent_predicted = pd.DataFrame(columns=['YearlyRent', 'Rooms'])

data_without_null = data_with_null.dropna()
#print(data_without_null.dtypes)
train_data_x = data_without_null.iloc[:,:5]
#print(train_data_x)
train_data_y = data_without_null.iloc[:,5]
#print(train_data_y)

#Training data
linreg.fit(train_data_x,train_data_y)

data_with_null = df_unclean[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent']].fillna(0)
test_data = data_with_null.iloc[:,:5]
#print(test_data)
rent_predicted['YearlyRent'] = pd.Series(linreg.predict(test_data))
#print(rent_predicted['YearlyRent'])
rent_predicted.YearlyRent = np.round(rent_predicted.YearlyRent)
df.YearlyRent.fillna(rent_predicted.YearlyRent,inplace=True)
#print(df.head(10))


'''
df_unclean_rooms_with_null = df_unclean[['SoldPrice','AskPrice','Size','MonthlyRent','YearlyRent','Rooms']].dropna()
df_unclean_rooms_without_null = df_unclean_rooms_with_null.dropna()
train_rooms_x = df_unclean_rooms_without_null.iloc[:,:5]
print(train_rooms_x)
train_rooms_y = df_unclean_rooms_without_null.iloc[:,5]
print(train_rooms_y)

#Training data for rooms
linreg.fit(train_rooms_x,train_rooms_y)
df_unclean_rooms_with_null = df_unclean[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent']].fillna(0)
test_room_data = df_unclean_rooms_with_null.iloc[:,:5]

rent_predicted['Rooms'] = pd.Series(linreg.predict(test_room_data))
print(rent_predicted['Rooms'])

'''

#The max and min range for number of rooms is very less compared to other features.
#So, it could be better to use either mean/median to impute empty fields.
room_median = df['Rooms'].median()
print(room_median)
df.Rooms.fillna(room_median, inplace=True)

df_unclean_rooms_with_null = df_unclean[['SoldPrice','AskPrice', 'Rooms','Size','MonthlyRent','YearlyRent','YearBuilt']].dropna()
df_unclean_rooms_without_null = df_unclean_rooms_with_null.dropna()
train_rooms_x = df_unclean_rooms_without_null.iloc[:,:6]
print(train_rooms_x)
train_rooms_y = df_unclean_rooms_without_null.iloc[:,6]
print(train_rooms_y)

#Training data for rooms
linreg.fit(train_rooms_x,train_rooms_y)
df_unclean_rooms_with_null = df_unclean[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent','YearBuilt']].fillna(0)
test_room_data = df_unclean_rooms_with_null.iloc[:,:6]

rent_predicted['YearBuilt'] = pd.Series(linreg.predict(test_room_data))
print(rent_predicted['YearBuilt'])
print(rent_predicted.columns)

rent_predicted.YearBuilt = np.round(rent_predicted.YearBuilt)

df.YearBuilt.fillna(rent_predicted.YearBuilt,inplace=True)
print(df.isnull().sum())

df.to_csv("result-clean.csv", index=False)