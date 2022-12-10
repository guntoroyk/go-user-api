-- Adminer 4.8.1 MySQL 8.0.31 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE IF NOT EXISTS `go_user_api_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `go_user_api_db`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone` varchar(255),
  `role` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`id`, `name`, `username`, `password`, `phone`, `role`, `created_at`, `updated_at`) VALUES
(1,	'admin',	'admin',	'$2a$14$lN4HlrtYZ.x2oOXSYVOwkuODuYtyNViQqELmE6RvF0BxKUIOn1/5m',	'085678765444',	'admin',	'2022-12-10 07:25:34',	'2022-12-10 07:25:34'),
(2,	'user',	'user',	'$2a$14$5NYSkvWjy7lF/3vPu0eXPOfkLn8CKL6m4HmFric6dk3ASNP/z4LEW',	'085678765445',	'user',	'2022-12-10 07:25:49',	'2022-12-10 07:25:49');

-- 2022-12-10 00:26:11
-- 2022-12-05 16:12:13