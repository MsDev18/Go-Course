-- +migrate Up
-- please read this article to andrestand why use VARCHAR(191)
-- https://www.grouproo.com/blog/varchar-191#why-varchar-and-not-text
CREATE TABLE `permissions` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(191) NOT NULL UNIQUE ,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE `permissions`;