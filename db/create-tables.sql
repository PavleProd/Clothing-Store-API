DROP TABLE IF EXISTS products;
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

INSERT INTO products (name, category, size, gender, is_for_kids, price, quantity) VALUES ('Kaput 25', 'Elegantna odeca', 'XL', 'Unisex', TRUE, 199.09, 72);