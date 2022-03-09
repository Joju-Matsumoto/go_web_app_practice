DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS comments CASCADE;

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    content TEXT,
    author VARCHAR(255)
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content TEXT,
    author VARCHAR(255),
    post_id INTEGER REFERENCES posts(id)
);