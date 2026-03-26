CREATE TABLE roles (
  id   SERIAL PRIMARY KEY,
  name TEXT   UNIQUE  NOT NULL
);

CREATE TABLE users (
    id         SERIAL    PRIMARY KEY,
    email      TEXT      UNIQUE NOT NULL,
    password   TEXT      NOT NULL,
    role_id    INTEGER   REFERENCES roles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)