-- todo_app.account_login_histories definition

CREATE TABLE `account_login_histories` (
    `account_login_history_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `account_id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(191) DEFAULT NULL,
    `password` varchar(191) DEFAULT NULL,
    `token` varchar(191) DEFAULT NULL,
    `token_expire_at` datetime DEFAULT NULL,
    `login_at` datetime DEFAULT NULL,
    `login_out_at` datetime DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`account_login_history_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;