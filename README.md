# Banking micro-service

## Description
shiny micro-service based on protobuf. Service can be used to list account, payments
Also you can make some CRUD operations for accounts. 
Also yoy can make some payments between accounts, but be aware of sending different currency!   

## How to run
running `make run`
this also required database and migrations: `make compose-up && make test-migrations-up`

## How to test
unit tests: `make test`

integration tests (after running service): `make test-integration`

repos tests (after setting up database): `make test-repos`
this also required database and migrations: `make compose-up && make test-migrations-up`

## Linting
lint project: `make lint`

linting rules located at ./.lint.yaml

## DOCS

Documentation is located at ./docs folder. 
Also you can open ./api folder to see description at `.proto` files.
Also you can open `http://localhost:8000` to see swagger.