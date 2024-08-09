CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name    VARCHAR(40) NOT NULL,
    email   VARCHAR(40) NOT NULL,
    password VARCHAR(100) NOT NULL
)