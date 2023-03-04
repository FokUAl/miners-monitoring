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
    miner_type VARCHAR(255) NOT NULL,
    shelf INTEGER DEFAULT 0,
    _row INTEGER DEFAULT 0,
    col INTEGER DEFAULT 0,
    owner_ VARCHAR(255) NOT NULL DEFAULT '',
    miner_status VARCHAR(255) NOT NULL,
    coin VARCHAR(255) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    mac_address VARCHAR(255) NOT NULL
);

CREATE TABLE miner_characteristics
(
    id SERIAL PRIMARY KEY,
    elapsed INTEGER DEFAULT 0,
    mhs_av FLOAT DEFAULT 0.00,
    temperature FLOAT DEFAULT 0.00,
    fan_speed_in INTEGER DEFAULT 0,
    fan_speed_out INTEGER DEFAULT 0,
    power_mode VARCHAR(255) NOT NULL,
    chip_temp_min FLOAT DEFAULT 0.00,
    chip_temp_max FLOAT DEFAULT 0.00,
    chip_temp_avg FLOAT DEFAULT 0.00,
    creation_date DATE NOT NULL
);