CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role_id INT NOT NULL DEFAULT 1 REFERENCES roles(id) 
);

-- Seed
INSERT INTO users(email, password, role_id)
VALUES
    ('joh@mail.com', 'test', 1),
    ('doe@gmail.com', 'password', 2);