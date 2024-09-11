CREATE TABLE links
(
    id        SERIAL PRIMARY KEY,
    original  TEXT UNIQUE NOT NULL,
    shortened TEXT UNIQUE NOT NULL
);