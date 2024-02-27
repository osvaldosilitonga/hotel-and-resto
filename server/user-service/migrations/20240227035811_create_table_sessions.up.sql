CREATE TABLE IF NOT EXISTS sessions(
    refresh_token VARCHAR PRIMARY KEY NOT NULL,
    access_token VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    role_id INT NOT NULL,
    exp BIGINT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);