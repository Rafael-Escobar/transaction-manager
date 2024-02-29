# transaction-manager

    The Transaction Manager is a microservice for managing accounts and transaction creation. This service provides the following methods:
    - Create account
    - Get account
    - Create Transaction
  
## API Documentation

    To access the API documentation, you can either:
    - Navigate to the docs folder and open the Swagger files in a Swagger editor.
    - Run the project locally and access the Swagger documentation at (documentation)[http://localhost:8080/v1/swagger/index.html] `http://localhost:8080/v1/swagger/index.html`

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

## Testing on Development Environment

### Get App Info
```
curl --location 'http://localhost:8080/info'
```

### Create account
```
curl --location 'http://localhost:8080/accounts' \
--header 'Content-Type: application/json' \
--data '{
"document_number": "{document_number}"
}'
```

### Get Account
```
curl --location 'http://localhost:8080/accounts/{id}'
```

### Create Transaction
```
curl --location 'http://localhost:8080/transactions' \
--header 'Content-Type: application/json' \
--data '{
"account_id": {account_id},
"operation_type_id": {operation_type_id},
"amount": {amount}
}'
```