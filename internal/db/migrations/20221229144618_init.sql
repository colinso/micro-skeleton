-- +goose Up
-- +goose StatementBegin

-- Create item table
CREATE TABLE IF NOT EXISTS item (
  id varchar PRIMARY KEY,
  name varchar NOT NULL,
  image_url varchar,
  count integer NOT NULL,
  price decimal NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE item;
-- +goose StatementEnd
