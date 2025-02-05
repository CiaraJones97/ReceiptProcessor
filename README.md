# Receipt Processor Solution

## Description
Simple webservice that takes JSON receipts and calculates point totals. Data will not save in-memory on application end.

# Installation
[$ git clone https://go.googlesource.com/example](https://github.com/CiaraJones97/ReceiptProcessor.git)

'''
$ cd receipts-webservice
$ go run .
'''

# EndPoints

## POST /receipts/process

* Payload : JSON Receipt
* Response: JSON with a generated ID for the receipt
* Example Curl request :

'''
$ curl http://localhost:8080/receipts/process --include --header "Content-Type: application/json" --request "POST" --data "{\"retailer\": \"Aldi\", \"purchaseDate\": \"2024-01-01\", \"purchaseTime\": \"10:00\", \"items\": [{\"shortDescription\": \"Water 12PK\", \"price\": \"3.00\"}, {\"shortDescription\": \"Sparking Cider\", \"price\": \"15.30\"}, {\"shortDescription\": \"Cheese Dip\", \"price\": \"4.26\"}], \"total\": \"22.56\"}"
'''

## GET /receipts/{ID}/points

* ID : ID of the receipt record to get the points for
* Data Type : string
* Response : JSON containing points calucated for the receipt
* Example Curl request :

'''
$ curl http://localhost:8080/receipts/82989cf5-99ea-4e4c-a6cd-d57a706dfbe3/points"
'''
