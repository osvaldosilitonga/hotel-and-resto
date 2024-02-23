CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(20) NOT NULL,
    role INT NOT NULL
);

-- Seed
INSERT INTO users(email, password, role)
VALUES
    ('joh@mail.com', 'test', '2'),
    ('doe@gmail.com', 'password', '1');