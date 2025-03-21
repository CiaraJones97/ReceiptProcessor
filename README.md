# Receipt Processor Solution

# Description

Simple webservice that takes JSON receipts and calculates point totals. Data will not save in-memory on application end.

## Clone the project

```cmd
$ git clone https://github.com/CiaraJones97/ReceiptProcessor.git
```

## Run the project

```cmd
$ cd .\ReceiptProcessor\receipts-webservice
$ go run .
```

This will spin up the local webserver to `GET` and `POST` against. In a seperate terminal instance, you can continue the following:

# Endpoints

## POST /receipts/process

### URL Format

`curl http://localhost:8080/receipts/process`

### Request Body

#### Fields
- `retailer`
    - Type: String
    - Should be name of the business, example: `Aldi`
- `purchaseDate`
    - Type: String
    - Should be made in `YYYY-MM-DD` format, example: `2024-01-01`
- `purchaseTime`
    - Type: String
    - Should be made in 24 Hour format `HH:mm`, example: `10:00`
- `items`
    - Type: Array of Objects
    - Should contain `shortDescription` and `price`
        - `shortDescription`
            - Type: String
            - Should be brief description of the item purchased, example: `Water 12PK`
        - `price`
            - Type: String
            - Should be made as 0.00, example: `4.26`
- `total`
    - Type: String
    - Total amount of the receipt, example: `22.56`

### Example JSON

```json
{
    "retailer": "Aldi",
    "purchaseDate": "2024-01-01",
    "purchaseTime": "10:00",
    "items": [
        {
            "shortDescription": "Water 12PK",
            "price": "3.00"
        },
        {
            "shortDescription": "Sparking Cider",
            "price": "15.30"
        },
        {
            "shortDescription": "Cheese Dip",
            "price": "4.26"
        }
    ],
    "total": "22.56"
}
```

## Response

The endpoint will respond back with a status code indicating success or failure.

### Example Curl Request:

```cmd
$ curl http://localhost:8080/receipts/process --include \
--header "Content-Type: application/json" \
--request "POST" \
--data "{
    \"retailer\": \"Aldi\",
    \"purchaseDate\": \"2024-01-01\",
    \"purchaseTime\": \"10:00\",
    \"items\": [
        {
            \"shortDescription\": \"Water 12PK\",
            \"price\": \"3.00\"
        },{
            \"shortDescription\": \"Sparking Cider\",
            \"price\": \"15.30\"
        },{
            \"shortDescription\": \"Cheese Dip\",
            \"price\": \"4.26\"
        }
    ],
    \"total\": \"22.56\"
}"
```

### Response

Returns a JSON of the ID.

```cmd
{"id":"ff81ee4d-e17d-4411-bf3e-5ebe9e352da9"}
```

### Status Codes

* 201 - Created: Successfully created Receipt ID
* 400 - Bad Request: Failed to create Receipt ID, Receipt is Invalid

## GET /receipts/{ID}/points

### URL Format

`curl http://localhost:8080/receipts/{ID}/process`

### Description

Fetches the amount of points a given Receipt ID is worth.

### Example Curl Request:


```cmd
$ curl http://localhost:8080/receipts/82989cf5-99ea-4e4c-a6cd-d57a706dfbe3/points"
```

### Response

Returns the Points of the Receipt ID

```cmd
{"points":28}
```

### Status Codes

* 202 - Accepted: The Request ID has been accepted.
* 404 - Not Found: The Request ID was not found.

## GET /receipts/query

### URL Format

`curl http://localhost:8080/receipts/query?...`

### Description

Fetches receipts that match a given criteria. If no paramaters are added, all receipts are returned.

### Parameters
- `points_greater_than`
    - Type: Integer
    - Request receipts greater than or equal to a specified amount of points
- `purchased_after`
    - Type: Datetime
    - Should be made in `YYYY-MM-DD HH:mm` format, example: `2024-01-01 10:00`

### Example Curl Request:

```cmd
$ curl "http://localhost:8080/receipts/query?purchased_after=2023-01-01 10:00&points_greater_than=10"
```

### Response
Returns a JSON with the resulting receipts

```json
{
    "result": [
        {
            "resultsReturned":1,
            "receipts": [
                {
                    "id":"jp2n4jh6-e17d-4411-bf3e-87be9e352db6","retailer":"Walgreens",
                    "purchaseDate":"2024-03-20",
                    "purchaseTime":"10:30",
                    "items": [
                        {
                            "shortDescription":"Chips",
                            "price":"2.00"
                        },
                        {
                            "shortDescription":"Pepsi",
                            "price":"5.00"
                        },
                        {
                            "shortDescription":"Water",
                            "price":"1.50"
                        }
                    ],
                    "total":"8.00",
                    "points":13
                }
            ]
        }
    ]
}
```

### Status Codes

* 202 - Accepted: The query has been accepted
* 404 - Not Found: The query format is invalid
