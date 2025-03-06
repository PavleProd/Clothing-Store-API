# Clothing Store API

RESTful API for online clothing store, including web-server implementation in **GO** and **PostgreSQL** database

## Deployment

To deploy server locally you need to do the following:
1. Have [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running
2. Clone the repository
3. Open Powershell and navigate to repository root folder
4. Run following command: `docker compose up --build`

- You can access API on port 8080
- Database will be already initialized with data from `db/init/products_init.csv`

## Components

### Requirements
- GO 1.24.0
- PostgreSQL 17
  
### Web-Server

Goal of this project was to write web-server mostly in pure GO. Some of the implemented functionalities:

- Converter from HTTP request to any data model, including HTTP request validator
- Simplified ORM with query builders with
- Prepared Statements for PostgreSQL to prevent SQL Injection

## API

### /api/v1/products

Resource for all clothing products. Search can go by any combination of the resource parameters

Available requests:
1. **GET**: query clothing products with filters (needs minimum User authorization)
2. **POST**: add new clothing product (needs minimum Admin authorization)

Model parameters:
- "name": string
- "category": string
- "size": \["S", "M", "L", "XL", "XXL"]
- "gender": \['Male', 'Female', 'Unisex']
- "is_for_kids": bool
- "price": decimal, in EUR
- "quantity": unsigned int

### /api/v1/login

Resource for authorization that works with [JWT](https://jwt.io/)

Available request:
1. **POST**: authorization that takes username and password and returns JWT token

Model parameters:
- "username": unique string
- "password": string
- "role": int (0 - user, 1 - admin)

## Using API

[PostMan](https://www.postman.com/downloads/) is recommended for testing this API

You can create your requests with API described above.
You can also import already created Postman Collection for testing this api in `testing/online_store_api.postman_collection.json` 

## Testing

You can execute all unit tests available by running `testing/execute_unit_tests.bat` 

## External Libraries

- [PostgreSQL Driver](https://github.com/lib/pq)
- [JWT Package](https://pkg.go.dev/github.com/golang-jwt/jwt/v5@v5.2.1)
