CREATE TABLE IF NOT EXISTS roles(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Seed
INSERT INTO roles (name)
VALUES
    ('admin'),
    ('user');