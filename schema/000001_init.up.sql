CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    role_ VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE miner_devices
(
    id SERIAL PRIMARY KEY,
    miner_type VARCHAR(255) NOT NULL DEFAULT ' ',
    shelf INTEGER DEFAULT 0,
    _row INTEGER DEFAULT 0,
    col INTEGER DEFAULT 0,
    owner_ VARCHAR(255) NOT NULL DEFAULT ' ',
    miner_status VARCHAR(255) NOT NULL DEFAULT ' ',
    coin VARCHAR(255) NOT NULL DEFAULT ' ',
    ip_address VARCHAR(255) UNIQUE,
    mac_address VARCHAR(255) DEFAULT ' '
);

CREATE TABLE miner_characteristics
(
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(255) NOT NULL,
    elapsed INTEGER DEFAULT 0,
    mhs_av FLOAT DEFAULT 0.00,
    temperature FLOAT DEFAULT 0.00,
    fan_speed_in INTEGER DEFAULT 0,
    fan_speed_out INTEGER DEFAULT 0,
    power_mode VARCHAR(255) NOT NULL,
    chip_temp_min FLOAT DEFAULT 0.00,
    chip_temp_max FLOAT DEFAULT 0.00,
    chip_temp_avg FLOAT DEFAULT 0.00,
    creation_date TIMESTAMP NOT NULL,
    CONSTRAINT fk_ip_address
        FOREIGN KEY (ip_address)
            REFERENCES miner_devices(ip_address)
            ON DELETE CASCADE
);

CREATE TABLE comments
(
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    comment VARCHAR(255)
);