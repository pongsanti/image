-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE images (
  id serial PRIMARY KEY,
  created_at timestamp with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  filename text NOT NULL DEFAULT '',
  href text NOT NULL DEFAULT '',
  file_type text NOT NULL DEFAULT '',
  file_size integer NOT NULL DEFAULT 0,
  creator_id integer NOT NULL DEFAULT 0
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE images;
