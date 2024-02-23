CREATE TABLE IF NOT EXISTS user_details(
    user_id UUID NOT NULL REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    birth DATE NOT NULL,
    address VARCHAR,
    gender VARCHAR(15) NOT NULL,
    created_at BIGINT,
    updated_at BIGINT
);