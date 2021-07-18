CREATE TABLE IF NOT EXISTS user_tab (
    `user_id` BIGINT(20) AUTO_INCREMENT PRIMARY KEY,
    `status` INT(11) UNSIGNED NOT NULL,
    `auth_methods` TEXT,
    `profile` TEXT,
    `meta_data` TEXT,
    `ctime` INT(11) UNSIGNED NOT NULL,
    `mtime` INT(11) UNSIGNED NOT NULL
)  ENGINE=INNODB;
