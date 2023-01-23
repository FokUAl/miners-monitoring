CREATE TABLE users
{
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    position INT DEFAULT 0 NOT NULL
};

CREATE TABLE rooms
{
    id SERIAL PRIMARY KEY,
    roomname VARCHAR(255) NOT NULL,
    creator VARCHAR(255) NOT NULL,
    capacity INT NOT NULL,
};