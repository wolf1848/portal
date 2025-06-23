CREATE DATABASE portal;
CREATE USER taxi WITH PASSWORD 'password';
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO taxi;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO taxi;
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE users 
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    profession VARCHAR(50),
    phone VARCHAR(10),
    telegram VARCHAR(30),
    city VARCHAR(30),
    email VARCHAR(50) NOT NULL UNIQUE,
    email_verified BOOLEAN DEFAULT FALSE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    locked BOOLEAN DEFAULT FALSE
);

CREATE TRIGGER update_users_modified
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash VARCHAR(512) NOT NULL UNIQUE,  -- Хэш токена
    device_info TEXT,
    ip_address VARCHAR(45),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    revoked BOOLEAN DEFAULT FALSE,
    
    CONSTRAINT token_unique UNIQUE (user_id, token_hash)
);

CREATE TRIGGER update_users_modified
BEFORE UPDATE ON refresh_tokens
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();