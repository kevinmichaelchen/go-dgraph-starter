CREATE TABLE IF NOT EXISTS users(
    id varchar(20) PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL
);
