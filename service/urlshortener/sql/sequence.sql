CREATE TABLE `sequence` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  `stub` varchar(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_stub` (`stub`)
) ENGINE = MyISAM DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'unique index generator';

/*
- command for index generation
REPLACE INTO sequence (stub) VALUES ('a');


- read the latest index
SELECT LAST_INSERT_ID();


- sharding with odd-even servers
auto-increment-increment = 2
auto-increment-offset = 1

auto-increment-increment = 2
auto-increment-offset = 2
*/