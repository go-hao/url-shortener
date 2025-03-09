CREATE TABLE `url_map` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  `long_url` varchar(2048) NOT NULL DEFAULT '',
  `long_url_md5` char(32) NOT NULL DEFAULT '',
  `short_url` varchar(11) NOT NULL DEFAULT '',

  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_long_url_md5` (`long_url_md5`),
  UNIQUE KEY `udx_short_url` (`short_url`)

) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;