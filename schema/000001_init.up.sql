CREATE TABLE users
{
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
};

CREATE TABLE miner_devices
{
    id SERIAL PRIMARY KEY,
    miner_type VARCHAR(255) NOT NULL,
    area VARCHAR(255) NOT NULL,
    miner_status VARCHAR(255) NOT NULL,
    coin VARCHAR(255) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    mac_address VARCHAR(255) NOT NULL,
    _pool VARCHAR(255) NOT NULL
}