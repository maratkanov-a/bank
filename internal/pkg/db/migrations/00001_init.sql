-- +goose Up
-- TODO: more types???
CREATE TYPE currency_type_enum AS ENUM (
  'USD', 'EUR', 'RU'
  );

-- TODO: more types???
CREATE TYPE direction_type_enum AS ENUM (
  'incoming', 'outgoing'
  );

CREATE TABLE account
(
  id           BIGSERIAL PRIMARY KEY NOT NULL,
  name         TEXT                  NOT NULL,
  balance      INT       DEFAULT 0,
  currency     currency_type_enum    NOT NULL,
  is_available BOOLEAN   DEFAULT TRUE,
  created_at   TIMESTAMP DEFAULT now()
);

CREATE TABLE payment
(
  id           BIGSERIAL PRIMARY KEY          NOT NULL,
  amount       INT                            NULL,
  account_from BIGINT REFERENCES account (id) NOT NULL,
  account_to   BIGINT REFERENCES account (id) NOT NULL,
  direction    direction_type_enum            NOT NULL,
  created_at   TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE account;
DROP TABLE payment;

DROP TYPE currency_type_enum;
DROP TYPE direction_type_enum;