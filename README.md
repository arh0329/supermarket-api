# SuperMarket API
[![Go](https://github.com/arh0329/supermarket-api/actions/workflows/go.yml/badge.svg)](https://github.com/arh0329/supermarket-api/actions/workflows/go.yml)
[![ci](https://github.com/arh0329/supermarket-api/actions/workflows/main.yml/badge.svg)](https://github.com/arh0329/supermarket-api/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/arh0329/supermarket-api/branch/main/graph/badge.svg)](https://codecov.io/gh/arh0329/supermarket-api)

Basic CRUD API in golang

- The produce database is only a single, in memory array of data and supports the following operations: add, delete, and fetch
- Supports adding more than one new produce item at a time
- The produce includes name, produce code, and unit price
- The produce name is alphanumeric(with spaces) and case insensitive
- The produce codes are sixteen characters long, with dashes separating each four character group
- The produce codes are alphanumeric and case insensitive
- The produce codes must be unique
- The produce unit price is a number with up to 2 decimal places

## Routes

Available routes within the API

### GetAll

Get all the produce in the database. Return JSON array of produce.

Example:
- Call: GET - http://localhost:8000/produce
- Response:

```json
[
    {
        "name": "Lettuce",
        "produceCode": "A12T-4GH7-QPL9-3N4M",
        "unitPrice": 3.46
    },
    {
        "name": "Peach",
        "produceCode": "E5T6-9UI3-TH15-QR88",
        "unitPrice": 2.99
    },
    {
        "name": "Green Pepper",
        "produceCode": "YRT6-72AS-K736-L4AR",
        "unitPrice": 0.79
    },
    {
        "name": "Gala Apple",
        "produceCode": "TQ4C-VV6T-75ZX-1RMR",
        "unitPrice": 3.59
    }
]
```

### GetByProductCode

Get a produce item from the database. Accepts a url parameter of Produce Code that will identify the item to fetch.

Example:
- Call: GET - http://localhost:8000/produce/TQ4C-VV6T-75ZX-1RMR
- Response:

```json
{
    "name": "Gala Apple",
    "produceCode": "TQ4C-VV6T-75ZX-1RMR",
    "unitPrice": 3.59
}
```

### AddProduce

Add one or more new produce items to the database. Accepts JSON array input with the below parameters:

- Name - required
- Produce Code - required
- Unit Price - required

Example:
- Call: POST - http://localhost:8000/produce
- Request Body:

```json
[
{
    "name": "Tomato",
    "produceCode": "T84C-VV6T-75ZX-1RMR",
    "unitPrice": 3.46
}
]
```
- Response:

```json
{
    "added": [
        "Tomato"
    ],
    "errors": [],
    "message": "Item(s) added"
}
```

### DeleteProduce

Delete a produce item from the database. Accepts a url parameter of Produce Code that will identify the item to delete.

Example:
- Call: DELETE - http://localhost:8000/produce/TQ4C-VV6T-75ZX-1RMR
- Response:

```json
{
    "message": "Item deleted"
}
```

## Build and Run

To build the application use the included Docker file and run the following command

```bash
docker build -t arh0329/supermarket-api:test -f Dockerfile .
```

Run the application with the following command:

```bash
docker run -it --rm -p 8000:8000 supermarket-api:test
```

This will run the application on port `8000`. Visit http://localhost:8000 to make requests.

The image can also be build from DockerHub:

```bash
docker pull arh0329/supermarket-api:latest
```

## Postman Tests

A Postman [collection](./tests/postman/Supermarket-API.postman_collection.json) is available to test functionality of the API. If running in Postman be sure to use the [Local](./tests/postman/Local.postman_environment.json) environment. Run collection in order for all tests to pass, as some test data is based on previous calls.

A [docker-compose.yaml](./docker-compose.yml) is also included to quickly run and test the API. The docker-compose file will build and run the API. It will also start a [newman](https://learning.postman.com/docs/running-collections/using-newman-cli/newman-with-docker/) container and run postman collection against the API

Use the following command to run the docker-compose

```bash
docker-compose up
```

Run the following command to remove all images/containers associated with the docker-compose

```bash
docker-compose down -v --rmi all
```