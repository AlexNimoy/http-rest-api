# HTTP Go API

Simple REST API on Golang with logs, CORS and Authentication middlewares.

## .env example

```
TEST_DATABASE_URL=host=localhost dbname=restapi_dev sslmode=disable user=developer password=developer port=5453
```

## Run PostgreSQL DB

```
docker-compose up
```

## Migrations

https://github.com/golang-migrate/migrate

```
migrate -path migrations -database "postgresql://developer:developer@localhost:5453/restapi_test?sslmode=disable" u
```

## Tests

```
make test
```

## Build && Run application

```
make && ./apiserver
```
