# Clothing Store API

Simple RESTful API for online clothing store. 

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

## Components

### Requirements
- GO 1.24.0
- PostgreSQL 17
  
### Web-Server

GO 1.24.0

Web-Server was implemented using pure GO. Only external library is a driver for PostgreSQL mentioned below.
Some of the implemented functionalities:

- Converter from HTTP request to any data model, including request validator
- Automatic SQL query builder from data model

## External Libraries

- [PostgreSQL Driver](https://github.com/lib/pq)
