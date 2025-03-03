# Clothing Store API

RESTful API for online clothing store, including web-server implementation in **GO** and **PostgreSQL** database

## API

### /api/v1/products

Resource for all clothing products. Search can go by any combination of the resource parameters

Model parameters:
- "name": string
- "category": string
- "size": \["S", "M", "L", "XL", "XXL"]
- "gender": \['Male', 'Female', 'Unisex']
- "is_for_kids": bool
- "price": decimal, in EUR
- "quantity": unsigned int

GET Query example:
`/api/v1/products?category=Sweater&is_for_kids=false&gender=Male`

## Deployment

To deploy server locally you need to do the following:
1. Download the repository
2. Open Powershell and navigate to repository root folder
3. Run following command: `docker compose up --build`
4. You can access API on port 8080, database will be already initialized with data from `db/init/products_init.csv`

## Components

### Requirements
- GO 1.24.0
- PostgreSQL 17
  
### Web-Server

Web-Server was implemented using pure GO. Only external library is a driver for PostgreSQL mentioned below.
Some of the implemented functionalities:

- Converter from HTTP request to any data model, including HTTP request validator
- Automatic SQL query builder from data model

## External Libraries

- [PostgreSQL Driver](https://github.com/lib/pq)
