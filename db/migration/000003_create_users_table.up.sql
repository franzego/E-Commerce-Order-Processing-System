CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar(50) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password_hash text NOT NULL,
    created_at timestamptz DEFAULT now()
);