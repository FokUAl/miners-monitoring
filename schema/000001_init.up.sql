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
    ip_address VARCHAR(255) UNIQUE,
    mac_address VARCHAR(255) DEFAULT ' '
);

CREATE TABLE miner_characteristics
(
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(255) NOT NULL,
    elapsed INTEGER DEFAULT 0,
    ths_av FLOAT DEFAULT 0.00,
    temperature FLOAT DEFAULT 0.00,
    fan_speed1 INTEGER DEFAULT 0,
    fan_speed2 INTEGER DEFAULT 0,
    fan_speed3 INTEGER DEFAULT 0,
    fan_speed4 INTEGER DEFAULT 0,
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

CREATE TABLE mac_ip
(
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(255) NOT NULL,
    mac_address VARCHAR(255) DEFAULT ' '
);

CREATE TABLE chat_history
(
    creation_date TIMESTAMP NOT NULL,
    content VARCHAR(255) NOT NULL,
    sender VARCHAR(255) NOT NULL,
    sender_role VARCHAR(255) NOT NULL,
    recipient VARCHAR(255),
    is_read BOOLEAN DEFAULT FALSE
);