# Go Minimal API

This is a minimal API example using Go and MySQL.

## Requirements

- Go 1.20
- Docker
- Docker Compose

## Installation

- Clone this repository
- Install dependencies: `go mod tidy`
- Run the database: `docker-compose up -d` (check architecture of your processor to use the correct image)
- Run the API: `go run main.go`

## Usage

Sample requests:

```bash
curl --location 'http://localhost:8080/products'
```

```bash
curl --location 'http://localhost:8080/products' \
--header 'Content-Type: application/json' \
--data '{    
    "name": "Product 1",
    "price": 14
}'
```