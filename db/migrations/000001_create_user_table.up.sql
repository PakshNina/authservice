CREATE TABLE users (
      id serial PRIMARY KEY,
      username VARCHAR(100) UNIQUE NOT NULL,
      password_hash CHAR(64) NOT NULL
);
-- Creating initial user for test purpose only
INSERT INTO users (username, password_hash) values ('user', '$2a$14$I8EjWVFx9k5zyogucf4b7ePRofKmG0ioqb0hbvrU/AgkntVk8v7v6')