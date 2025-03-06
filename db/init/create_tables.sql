DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS Size;
DROP TYPE IF EXISTS Gender;

CREATE TYPE Size AS ENUM ('XS', 'S', 'M', 'L', 'XL', 'XXL');
CREATE TYPE Gender AS ENUM ('Male', 'Female', 'Unisex');

CREATE TABLE products (
    id  SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) NOT NULL,
    size Size NOT NULL,
    gender Gender NOT NULL,
    is_for_kids BOOLEAN NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role INT NOT NULL
);

COPY products(name, category, size, gender, is_for_kids, price, quantity)
FROM '/docker-entrypoint-initdb.d/products_init.csv'
DELIMITER ','
CSV HEADER;

COPY users(username, password, role)
FROM '/docker-entrypoint-initdb.d/users_init.csv'
DELIMITER ','
CSV HEADER;