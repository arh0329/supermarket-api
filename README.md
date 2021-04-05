# SuperMarket API
[![Actions Status](https://github.com/arh0329/supermarket-api/workflows/Go/badge.svg)](https://github.com/arh0329/supermarket-api/actions)
[![codecov](https://codecov.io/gh/arh0329/supermarket-api/branch/main/graph/badge.svg)](https://codecov.io/gh/arh0329/supermarket-api)

Basic CRUD API in golang

- The produce database is only a single, in memory array of data and supports the following operations: add, delete, and fetch
- Supports adding more than one new produce item at a time
- The produce includes name, produce code, and unit price
- The produce name is alphanumeric and case insensitive
- The produce codes are sixteen characters long, with dashes separating each four character group
- The produce codes are alphanumeric and case insensitive
- The produce unit price is a number with up to 2 decimal places

## Routes

Available routes within the API

### GetAll

Get all the produce in the database. Return JSON array of produce.

### GetByProductCode

Get a produce item from the database. Accepts a url parameter of Produce Code that will identify the item to delete.

### AddProduce

Add one or more new produce items to the database. Accepts JSON input with the following parameters:

- Name - required
- Produce Code - required
- Unit Price - required

### DeleteProduce

Delete a produce item from the database. Accepts a url parameter of Produce Code that will identify the item to delete.

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