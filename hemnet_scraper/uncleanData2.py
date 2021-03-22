import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression
import seaborn as sns
import matplotlib.pyplot as plt
import random
from pylab import rcParams

linreg = LinearRegression()

df = pd.read_csv("result-clean-2.csv")
duplicate_df = df[df.duplicated()]
print(duplicate_df)
print(df.shape)
print(df.isnull().sum())


df_yearlyrent = df['YearlyRent'].count()
#print(df_yearlyrent)
df_yearlyrent_empty = df['YearlyRent'].isnull().sum()
#print(df_yearlyrent_empty)
df_yearlyrent_sum = df_yearlyrent + df_yearlyrent_empty
#print("Empty fields in yearlyRent column: {:.2f}%".format(df_yearlyrent_empty/df_yearlyrent_sum * 100))
'''':cvar
With more than 30% empty fields in yearlyRent it is wise to impute empty data using linear regression rather than dropping the rows. 
We should over come loss of data. let's impute yearlyRent. 
'''
data_with_null = df[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent']].dropna()

#print(data_with_null.isnull().sum())
predictions = pd.DataFrame(columns=['yearlyRent'])

data_without_null = data_with_null.dropna()

#print(data_without_null.dtypes)
train_data_x = data_without_null.iloc[:,:5]
#print(train_data_x)
train_data_y = data_without_null.iloc[:,5]
#print(train_data_y)

#Training data
linreg.fit(train_data_x,train_data_y)

data_with_null = df[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent']].fillna(0)
test_data = data_with_null[['SoldPrice', 'AskPrice','Rooms','Size','MonthlyRent']]
#print(test_data)
predictions['YearlyRent'] = pd.Series(linreg.predict(test_data))

#print(predictions['YearlyRent'].head(20))
predictions.YearlyRent = np.round(predictions.YearlyRent)
df.YearlyRent.fillna(predictions.YearlyRent,inplace=True)
#print(df['YearlyRent'].head(20))

#df.to_csv("reventh_clean.csv", index=False)

#print(df.isnull().sum())

df_rooms = df['Rooms'].count()
#print(df_rooms)

df_rooms_empty = df['Rooms'].isnull().sum()
#print(df_rooms_empty)

df_rooms_sum = df_rooms + df_rooms_empty
#print(df_rooms_sum)
print("Empty fields in Rooms column: {:.2f}%".format(df_rooms_empty/df_rooms_sum * 100))

''':cvar
Empty fields in Rooms column is 6.80% Similaryly empty fields in Size column is further less.
We can either impute missing fields by either linear regression or by median values. Let's give a try   
'''

rooms_median = df['Rooms'].median()
#print(rooms_median)
df.Rooms.fillna(rooms_median, inplace=True)

size_median = df['Size'].median()
#print(size_median)
df.Size.fillna(size_median, inplace=True)
print(df.isnull().sum())
#df.to_csv("reventh_clean.csv", index=False)

monthly_data_with_null = df[['SoldPrice','AskPrice','Rooms','Size','YearlyRent','MonthlyRent']].dropna()

monthly_data_without_null = monthly_data_with_null.dropna()

monthly_train_data_x = monthly_data_without_null.iloc[:,:5]
monthly_train_data_y = monthly_data_without_null.iloc[:,5]

linreg.fit(monthly_train_data_x,monthly_train_data_y)

monthly_data_with_null = df[['SoldPrice','AskPrice','Rooms','Size','YearlyRent','MonthlyRent']].fillna(0)

monthly_test_data = monthly_data_with_null[['SoldPrice','AskPrice','Rooms','Size','YearlyRent']]

predictions['MonthlyRent'] = pd.Series(linreg.predict(monthly_test_data))
print(predictions['MonthlyRent'])

predictions.MonthlyRent = np.round(predictions.MonthlyRent)
df.MonthlyRent.fillna(predictions.MonthlyRent,inplace=True)
#print(df.isnull().sum())
#df.to_csv("reventh_clean.csv", index=False)

'''':cvar
Impute AskPrice
'''

askprice_with_null = df[['SoldPrice','Rooms','Size','YearlyRent','MonthlyRent','AskPrice']].dropna()

askprice_without_null = askprice_with_null.dropna()

askprice_train_data_x = askprice_without_null.iloc[:,:5]
askprice_train_data_y = askprice_without_null.iloc[:,5]

linreg.fit(askprice_train_data_x,askprice_train_data_y)
askprice_with_null = df[['SoldPrice','Rooms','Size','YearlyRent','MonthlyRent','AskPrice']].fillna(0)

askprice_test_data = askprice_with_null[['SoldPrice','Rooms','Size','YearlyRent','MonthlyRent']]
predictions['AskPrice'] = pd.Series(linreg.predict(askprice_test_data))
#print(predictions['AskPrice'])
predictions.AskPrice = np.round(predictions.AskPrice)
df.AskPrice.fillna(predictions.AskPrice,inplace=True)



'''':cvar
Impute yearBuilt
'''
yearbuilt_with_null = df[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent','YearBuilt']].dropna()
yearbuilt_without_null = yearbuilt_with_null.dropna()

yearbuilt_train_data_x = yearbuilt_without_null.iloc[:,:6]
yearbuilt_train_data_y = yearbuilt_without_null.iloc[:,6]
#print(yearbuilt_train_data_x)
#print(yearbuilt_train_data_y)

linreg.fit(yearbuilt_train_data_x,yearbuilt_train_data_y)
yearbuilt_with_null = df[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent','YearBuilt']].fillna(0)

yearbult_test_data = yearbuilt_with_null[['SoldPrice','AskPrice','Rooms','Size','MonthlyRent','YearlyRent']]
predictions['YearBuilt'] = pd.Series(linreg.predict(yearbult_test_data))
#print(predictions['YearBuilt'])

predictions.YearBuilt = np.round(predictions.YearBuilt)
df.YearBuilt.fillna(predictions.YearBuilt, inplace=True)

print(df.isnull().sum())

df.to_csv("reventh_clean.csv", index=False)

