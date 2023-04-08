-- todo_app.tasks definition

CREATE TABLE `tasks` (
  `task_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `title` longtext NOT NULL,
  `description` longtext NOT NULL,
  `completed` tinyint(1) NOT NULL,
  `task_status` varchar(191) DEFAULT NULL,
  `completed_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`task_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;