-- +migrate Up
-- MYSQL 8.0 set the role value to `user` for all records
-- Don't change the order of enum values
-- TODO - find a better solution instead of keeping order !!!
ALTER TABLE `users` ADD COLUMN `role` ENUM("user","admin") NOT NULL;
-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;