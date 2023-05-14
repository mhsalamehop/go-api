DROP TABLE users_table;
DROP TABLE saved_movies;
CREATE TABLE IF NOT EXISTS saved_movies (
    id SERIAL PRIMARY KEY,
    title text,
    overview text
);
CREATE TABLE IF NOT EXISTS users_table (
    id SERIAL PRIMARY KEY,
    email text UNIQUE,
    username text UNIQUE,
    password text,
    token text,
    role text DEFAULT 'user'
);
-- CREATE INDEX ON "users_table" ("token");