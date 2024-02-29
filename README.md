# Transaction Manager

The Transaction Manager is a microservice for managing accounts and transaction creation. This service provides the following methods:

- Create Account
- Get Account
- Create Transaction
  
## API Documentation

To access the API documentation, you can either:

- Navigate to the docs folder and open the Swagger files in a Swagger editor.
- Run the project locally and access the Swagger documentation at [documentation](http://localhost:8080/v1/swagger/index.html) `http://localhost:8080/v1/swagger/index.html`

## Requirements

To run this project some system requirements are needed:

- Golang installed (version 1.20.8) or ASDF package manager
- Docker
- DockerCompose

## How run the project

1. Clone de repo
2. Install tools `make install-tools` if you already have golang installed just run `make install-go-tools`
3. Raise de database `make dev/up`
4. Run migrations `ENV=dev make migrate-up`
5. Run locally the application `make run-local`

### How to Update Mocks

To update mocks, run the following command:
    ```make mock-all```

### How to Update Swagger Documentation

To update Swagger documentation, run:
    ```make doc/update```

### Other commands

Run uni tests
    ```make test```

To check lint
    ```make lint```

To check security vulnerability
    ```make security-check```

To create migration
    ```make migrate-create```

To downgrade migration
    ```make migrate-down```

## Testing on Development Environment

### Get App Info

```bash
curl --location 'http://localhost:8080/info'
```

### Create account

```bash
curl --location 'http://localhost:8080/accounts' \
--header 'Content-Type: application/json' \
--data '{
"document_number": "{document_number}"
}'
```

### Get Account

```bash
curl --location 'http://localhost:8080/accounts/{id}'
```

### Create Transaction

```bash
curl --location 'http://localhost:8080/transactions' \
--header 'Content-Type: application/json' \
--data '{
"account_id": {account_id},
"operation_type_id": {operation_type_id},
"amount": {amount}
}'
```

## Next steps

Future changes and improvements mapped:

- [ ] Add container for application on docker-compose file
- [ ] Improve test coverage
- [ ] Add [air](https://github.com/cosmtrek/air) live reload
- [ ] Create wrapper for package log
- [ ] Add container DI to the project

