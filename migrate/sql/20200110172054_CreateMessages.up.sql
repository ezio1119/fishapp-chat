CREATE TABLE `messages`(
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `body` VARCHAR(255) NOT NULL,
  `room_id` INT(11) NOT NULL,
  `user_id` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`room_id`) 
    REFERENCES rooms(`id`)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_ja_0900_as_cs;