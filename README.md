# Golang REST API

RESTful API with Go and PostgreSQL using Docker.

## Preparing to launch a project

- Install [Golang](https://golang.org/doc/install)
- Install [Docker](https://docs.docker.com/get-docker/)
- Install [Migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
- Install Postgres:

```bash
docker pull postgres:15-alpine
```

## Install & Run

```bash
# download this project
go get github.com/esmira23/go-postgres-docker
```

Before running API server, you should build docker image and run it. Also you need to create database and use db/migration to create table.

```bash
# build and run docker-compose
make build && make run
# create db and migrate table
make createdb && make migrateup
```

Now you can build and run project.

```bash
# build and run
cd go-postgresql-docker
go build
./go-postgresql-docker

# API Endpoint : http://127.0.0.1:8000
```

## Project tree

```
├── db/migration            // migration for database
│   ├── 000001_init_schema.down.sql
│   └── 000001_init_schema.up.sql
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── config
│   └── config.go           // Configuration for database
├── controllers
│   └── controller.go       // Controllers for application
├── csvparser
│   ├── csvparser.go        // CSV parser
│   └── example.csv         // csv data
├── routers
│   └── router.go           // Routers for application
├── models
│   └── model.go            // Models for application
├── Makefile
├── docker-compose.yml      // docker compose file
├── main_test.go            // Unit tests
└── main.go
```

## API

| Method | Path                                 | Description                                 |
| ------ | ------------------------------------ | ------------------------------------------- |
| `GET`  | /api/all                             | Get all data                                |
| `GET`  | /api/transaction/:id                 | Get data by transaction id                  |
| `GET`  | /api/terminal                        | Get data by terminal id (allow several ids) |
| `GET`  | /api/status/:status                  | Get data by status (accepted/declined)      |
| `GET`  | /api/payment_type/:type              | Get data by payment type (cash/card)        |
| `GET`  | /api/date/?from=date_from&to=date_to | Get data by date_post by periods (from/to)  |
| `GET`  | /api/payment_narrative/:narrative    | Get data by partially payment_narrative     |

## {•••} Swagger API

Go to API Docs page: http://127.0.0.1:8000/swagger/index.html
![Swagger](docs/Swagger.png)

## Run all the Unit test cases

```bash
go test -v
```

## Tasks

- [x] Parse \*.csv file and save result to the database.
- [x] Filter and upload data in JSON format in the response:
  - search by transaction_id
  - search by terminal_id (possibility of specifying several ids)
  - search by status (accepted/declined)
  - search by payment_type (cash/card)
  - search by date_post by periods (from/to)
  - search on the partially specified payment_narrative
- [x] Use relational database: PostgreSQL.
- [x] API documentation.
- [x] Use docker and docker-compose.
