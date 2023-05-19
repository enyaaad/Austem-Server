CREATE TABLE IF NOT EXISTS users(
    id integer  NOT NULL,
    name text NOT NULL,
    password text NOT NULL
);
INSERT INTO users (id, name, password) VALUES (1,'asd','123'),(2,'zxc','1234')