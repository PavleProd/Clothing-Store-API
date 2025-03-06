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

Web-Server was implemented using pure GO. Only external library is a driver for PostgreSQL mentioned below.
Some of the implemented functionalities:

- Converter from HTTP request to any data model, including HTTP request validator
- Automatic SQL query builder from data model

## API

### /api/v1/products

Resource for all clothing products. Search can go by any combination of the resource parameters

Model parameters:
- "name": string
- "category": string
- "size": \["S", "M", "L", "XL", "XXL"]
- "gender": \['Male', 'Female', 'Unisex']
- "is_for_kids": bool (Default: false)
- "price": decimal, in EUR
- "quantity": unsigned int

### Using API

[PostMan](https://www.postman.com/downloads/) is recommended for testing this API

1. GET Query example:
URL: `localhost:8080/api/v1/products?category=Sweater&is_for_kids=false&gender=Male`

The Response will be either: 
- JSON array of the requested resource rows
- Appropriate Error Code with message  

2. POST Query example:

URL: `localhost:8080/api/v1/products`

Body:
```
{
    "category": "Shirts",
    "gender": "Male",
    "is_for_kids": false,
    "name": "Polo Shirt",
    "price": 64.24,
    "quantity": 3,
    "size": "S"
}
```

The Response will be either:
- StatusOK (200) if POST was successful
- Appropriate Error Code with message

## Testing

You can execute all unit tests available by running `execute_unit_tests.bat` 

## External Libraries

- [PostgreSQL Driver](https://github.com/lib/pq)
