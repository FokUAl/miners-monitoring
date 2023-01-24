CREATE TABLE users
{
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    position INT DEFAULT 0 NOT NULL
};
