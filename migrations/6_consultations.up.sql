CREATE TABLE IF NOT EXISTS `consultations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `mentor_id` bigint unsigned DEFAULT NULL,
  `desirable_date` longtext,
  `message` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_consultations_deleted_at` (`deleted_at`),
  KEY `fk_consultations_user` (`user_id`),
  KEY `fk_consultations_mentor` (`mentor_id`),
  CONSTRAINT `fk_consultations_mentor` FOREIGN KEY (`mentor_id`) REFERENCES `mentors` (`id`),
  CONSTRAINT `fk_consultations_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;