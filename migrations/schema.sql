CREATE DATABASE IF NOT EXISTS library;
USE library;

DROP TABLE IF EXISTS `library`;
CREATE TABLE IF NOT EXISTS `library` (
    `id` INT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL,
	`isbn` VARCHAR(255) NOT NULL,
	`year` VARCHAR(255) NOT NULL,
	`author` VARCHAR(255) NOT NULL,
	`description` TEXT,
	PRIMARY KEY (`id`)
);