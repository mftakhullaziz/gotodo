-- todo_app.user_details definition

CREATE TABLE `user_details` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `email` varchar(191) NOT NULL,
  `name` longtext NOT NULL,
  `mobile_phone` bigint(20) NOT NULL,
  `address` longtext NOT NULL,
  `status` longtext NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;