
-- +migrate Up

CREATE TABLE USERS (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  password VARCHAR(255) NOT NULL,
  modified_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL
);

-- +migrate Down

DROP TABLE USERS;
