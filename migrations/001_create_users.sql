--CREATE TABLE USERS

CREATE TYPE department_type AS ENUM ('IT','product', 'hr', 'sales')

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email  TEXT NOT NULL UNIQUE,
    department department_type NOT NULL DEFAULT 'IT',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


