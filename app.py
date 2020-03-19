#!/bin/python

import requests

url = "http://localhost:8080/v1/example/echo"
headers = { 'Content-Type': 'application/json'}
data= { 'value': 'example-value'}

# data =/= json
res = requests.post(url=url, headers=headers, json=data)
print("Response Status Code: {}".format(res.status_code))
print(res.json())
