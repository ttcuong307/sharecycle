-- +goose Up
CREATE TABLE users (
  id varchar(255) PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "users";