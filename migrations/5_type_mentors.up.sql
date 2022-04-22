CREATE TABLE IF NOT EXISTS `type_mentors` (
  `mentor_id` bigint unsigned NOT NULL,
  `type_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`mentor_id`,`type_id`),
  KEY `fk_type_mentors_type` (`type_id`),
  CONSTRAINT `fk_type_mentors_mentor` FOREIGN KEY (`mentor_id`) REFERENCES `mentors` (`id`),
  CONSTRAINT `fk_type_mentors_type` FOREIGN KEY (`type_id`) REFERENCES `types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;