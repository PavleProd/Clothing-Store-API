# Clothing Store API

Simple RESTful API for online clothing store. 

## API

### /api/v1/products

- Resource for all clothing products

Model parameters:
- "name" - string
- "category" - string
- "size" - "S", "M", "L", "XL", "XXL"
- "price" - float32, in EUR
- "quantity" - uint

## Tech Stack

Web-Server:

GO 1.24.0

- Whole backend (TCP communication, Data Proccessing) implemented using pure GO
