-- +goose Up
CREATE TABLE account (

);

CREATE TABLE payment (

);

-- +goose Down
DROP TABLE account;
DROP TABLE payment;
