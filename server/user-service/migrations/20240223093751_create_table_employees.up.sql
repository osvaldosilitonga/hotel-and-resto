CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS employees(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(25) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role_id int REFERENCES roles(id)
);

-- Seed
INSERT INTO employees(username, password, role_id)
VALUES
    ('john.doe', 'password', 1),
    ('jane.smith', 'test123', 2);