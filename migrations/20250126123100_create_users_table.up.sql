CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL,
                       deleted_at TIMESTAMP WITH TIME ZONE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


-- CREATE TABLE users (
--                        id SERIAL PRIMARY KEY,
--                        username VARCHAR(255) NOT NULL UNIQUE,
--                        email VARCHAR(255) NOT NULL UNIQUE,
--                        password_hash VARCHAR(255) NOT NULL,
--                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );