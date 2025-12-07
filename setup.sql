CREATE DATABASE gorest;

\c gorest;

CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   first_name VARCHAR(100) NOT NULL,
   last_name VARCHAR(100) NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   phone_number VARCHAR(20)
);

INSERT INTO users (first_name, last_name, email, phone_number)
VALUES
    ('John', 'Doe', 'john.doe@example.com', '1234567890'),
    ('Jane', 'Smith', 'jane.smith@example.com', '0987654321');