-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NULL,
  `mmr` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_users_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "teams" table
CREATE TABLE `teams` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_one_id` bigint unsigned NULL,
  `user_two_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_teams_user_one` (`user_one_id`),
  INDEX `fk_teams_user_two` (`user_two_id`),
  INDEX `idx_teams_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_teams_user_one` FOREIGN KEY (`user_one_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_teams_user_two` FOREIGN KEY (`user_two_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "matches" table
CREATE TABLE `matches` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `team_one_id` bigint unsigned NULL,
  `team_two_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_matches_team_one` (`team_one_id`),
  INDEX `fk_matches_team_two` (`team_two_id`),
  INDEX `idx_matches_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_matches_team_one` FOREIGN KEY (`team_one_id`) REFERENCES `teams` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_matches_team_two` FOREIGN KEY (`team_two_id`) REFERENCES `teams` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
