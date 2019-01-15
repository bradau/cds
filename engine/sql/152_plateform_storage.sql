-- +migrate Up
ALTER TABLE platform_model RENAME COLUMN file_storage TO artifacts;
ALTER TABLE platform_model DROP COLUMN block_storage;


-- +migrate Down
ALTER TABLE platform_model RENAME COLUMN artifacts TO file_storage;
ALTER TABLE platform_model ADD COLUMN block_storage BOOLEAN default false;
