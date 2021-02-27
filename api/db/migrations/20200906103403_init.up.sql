CREATE TABLE IF NOT EXISTS todos(
    id varchar(20) PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title varchar(255) NOT NULL,
    done boolean NOT NULL
);
