CREATE TABLE IF NOT EXISTS `mentors` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext,
  `email` longtext,
  `image` longtext,
  `click_count` bigint DEFAULT NULL,
  `destination` longtext,
  `edu_org` longtext,
  `term_id` bigint unsigned DEFAULT NULL,
  `price` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_mentors_deleted_at` (`deleted_at`),
  KEY `fk_mentors_term` (`term_id`),
  CONSTRAINT `fk_mentors_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=x DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;