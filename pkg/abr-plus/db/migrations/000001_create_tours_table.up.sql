-- db/migrations/{timestamp}_create_tables.up.sql

CREATE TABLE IF NOT EXISTS "User" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "ProductItem" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price INT NOT NULL,
    image VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "Order" (
    id SERIAL PRIMARY KEY,
    user_id INT,
    product_item_id INT,
    quantity INT NOT NULL,
    total_price INT NOT NULL,
    order_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "User" (id),
    FOREIGN KEY (product_item_id) REFERENCES "ProductItem" (id)
);
