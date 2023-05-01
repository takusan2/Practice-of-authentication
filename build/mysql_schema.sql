
SELECT UUID();

SELECT * FROM users WHERE id = '57ef3f22-e7d0-11ed-a43d-3dece8c83097';

SELECT * FROM users;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `tests`;

CREATE TABLE `users` (
  `id` VARCHAR(36) PRIMARY KEY NOT NULL DEFAULT (UUID()),
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL UNIQUE,
  `password_hash` VARCHAR(255) NOT NULL,
  `salt` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

/* 
CREATE TABLE `folders` (
  `id` serial PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `parent_id` bigint,
  `user_id` uuid,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `closure` (
  `ancestor` bigint NOT NULL,
  `descendants` bigint NOT NULL,
  `user_id` uuid,
  PRIMARY KEY (`ancestor`, `descendants`)
);

CREATE TABLE `notes` (
  `id` serial PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `content` text,
  `folder_id` bigint,
  `user_id` uuid,
  `created_at` datetime,
  `updated_at` datetime
);

ALTER TABLE `folders` ADD FOREIGN KEY (`parent_id`) REFERENCES `folders` (`id`);

ALTER TABLE `folders` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `closure` ADD FOREIGN KEY (`ancestor`) REFERENCES `folders` (`id`);

ALTER TABLE `closure` ADD FOREIGN KEY (`descendants`) REFERENCES `folders` (`id`);

ALTER TABLE `closure` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `notes` ADD FOREIGN KEY (`folder_id`) REFERENCES `folders` (`id`);

ALTER TABLE `notes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`); */
