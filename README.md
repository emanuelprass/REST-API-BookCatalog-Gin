# REST-API-BookCatalog-Gin

This is an simple REST API. This API build with Go Language with gin framework.
The purpose of this API is to learn to make a REST API.

## Prerequisite

- Go version 1.16.5
- Run command go mod tidy to install dependency
- MySQl Database

## Project structure

```
REST-API-BookCatalog
└── config          # load any config at .env.example
└── entity          # seems like model
└── handler         # seems like controller
└── transport       # standarization request from client and response data
|   ├── request
|   ├── response
└── usecase         # manage apps rule/flow
└── repository      # manage any parameter from usecase into query in database
└── server          # routing directory
main.go
```

## How to run

- Clone this repository
- Please run `cp .env.example .env` at your terminal
- Add your environment variable in file `.env`
- Run go run main.go
- Run page localhost:8080/{endpoint} in your postman