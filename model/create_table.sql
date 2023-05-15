DROP TABLE IF EXISTS `user`
CREATE TABLE `user`
(
    `id`          BIGINT NOT NULL AUTO_INCREMENT,
    `user_id`     BIGINT NOT NULL,
    `username`    VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`    VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email`       VARCHAR(64) COLLATE utf8mb4_general_ci,
    `gender`      TINYINT NOT NULL DEFAULT '0',
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`
(
    `id`             INT NOT NULL AUTO_INCREMENT,
    `community_id`   INT UNSIGNED NOT NULL,
    `community_name` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction`   VARCHAR(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time`    TIMESTAMP                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`    TIMESTAMP                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community`
VALUES ('1', '1', 'Go', 'Golang', '2023-05-01 08:10:10', '2023-05-01 08:10:10');
INSERT INTO `community`
VALUES ('2', '2', 'Leetcode', '刷题刷题', '2023-05-01 08:20:10', '2023-05-01 08:20:10');
INSERT INTO `community`
VALUES ('3', '3', 'CS:GO', 'Rush B', '2023-05-01 08:30:10', '2023-05-01 08:30:10');
INSERT INTO `community`
VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟', '2023-05-01 08:40:10', '2023-05-01 08:40:10');

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`
(
    `id`           BIGINT NOT NULL AUTO_INCREMENT,
    `post_id`      BIGINT NOT NULL COMMENT '帖子id',
    `title`        VARCHAR(128)                             NOT NULL COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content`      VARCHAR(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_id`    BIGINT NOT NULL COMMENT '作者的用户id',
    `community_id` BIGINT NOT NULL COMMENT '所属社区',
    `status`       TINYINT NOT NULL DEFAULT '1' COMMENT '帖子状态',
    `create_time`  TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY            `idx_author_id` (`author_id`),
    KEY            `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;