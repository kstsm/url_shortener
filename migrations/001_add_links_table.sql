CREATE TABLE links
(
    id        SERIAL PRIMARY KEY,
    original  TEXT UNIQUE NOT NULL,
    shortened TEXT UNIQUE NOT NULL
);

CREATE TABLE tg_links
(
    link_id INT REFERENCES links (id) NOT NULL,
    tg_chat_id INT NOT NULL,
    title   TEXT,
    UNIQUE (link_id, tg_chat_id)
);
