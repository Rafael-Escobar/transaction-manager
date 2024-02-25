# transaction-manager
microservice for manage transactions


## How test on development environment

### Get App Info
```
curl --location 'http://localhost:8080/info'
```

### Create account
```
curl --location 'http://localhost:8080/accounts' \
--header 'Content-Type: application/json' \
--data '{
"document_number": "12345678900"
}'
```

### Get Account
```
curl --location 'http://localhost:8080/accounts/12345678900'
```

### Create Transaction
```
curl --location 'http://localhost:8080/transactions' \
--header 'Content-Type: application/json' \
--data '{
"account_id": 1,
"operation_type_id": 4,
"amount": 123.45
}'
```