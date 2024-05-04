-- Modify "users" table
ALTER TABLE `users` MODIFY COLUMN `name` varchar(191) NULL, ADD UNIQUE INDEX `uni_users_name` (`name`);
