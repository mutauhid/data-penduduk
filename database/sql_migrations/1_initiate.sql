-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE user (
    id VARCHAR(5) PRIMARY KEY,
    username VARCHAR(25),
    password VARCHAR(250)
);

CREATE TABLE province (
    id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Regency
CREATE TABLE regency (
    id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    province_id VARCHAR(5) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (province_id) REFERENCES province (id)
);

-- Tabel District
CREATE TABLE district (
    id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    regency_id VARCHAR(5) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (regency_id) REFERENCES regency (id)
);

-- Tabel People
CREATE TABLE people (
    id VARCHAR(5) PRIMARY KEY,
    nik VARCHAR(16) NOT NULL,
    name VARCHAR(50) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    dob DATE NOT NULL,
    pob VARCHAR(35) NOT NULL,
    province_id VARCHAR(5) NOT NULL,
    regency_id VARCHAR(5) NOT NULL,
    district_id VARCHAR(5) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (province_id) REFERENCES province (id),
    FOREIGN KEY (regency_id) REFERENCES regency (id),
    FOREIGN KEY (district_id) REFERENCES district (id)
);

-- +migrate StatementEnd