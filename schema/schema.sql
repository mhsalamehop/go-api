DROP TABLE users_table;
CREATE TABLE IF NOT EXISTS saved_movies
(
    id SERIAL PRIMARY KEY,
    title text,
    release_date text,
    poster_path text,
    overview text
);

CREATE TABLE IF NOT EXISTS users_table
(
    id SERIAL PRIMARY KEY,
    email text,
    username text,
    password text,
    token text
);