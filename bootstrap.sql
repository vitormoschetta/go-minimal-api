DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id VARCHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO
    products (id, `name`, price)
VALUES
    ('5c2e6cd9-30f5-4714-98de-f4e1139b817c', 'Product 1', 10.00),
    ('e3c08c86-8046-474d-9b88-3786dbbdd226', 'Product 2', 20.00),
    ('0327279b-f60d-485c-b6f5-1a7cd39981c9', 'Product 3', 30.00);