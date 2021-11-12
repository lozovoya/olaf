CREATE TABLE groups
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    isActive BOOLEAN DEFAULT false NOT NULL,
    created  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    password VARCHAR(200) NOT NULL,
    email VARCHAR (100) NOT NULL UNIQUE,
    isActive BOOLEAN DEFAULT false NOT NULL,
    group_id BIGINT NOT NULL REFERENCES groups,
    created  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


