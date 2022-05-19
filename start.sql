CREATE DATABASE IF NOT EXISTS `go_project_example` ;

USE `go_project_example`;

DROP TABLE IF EXISTS `topics`;

CREATE TABLE `topics` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `title` varchar(128) NOT NULL default '' COMMENT '标题',
    `content` text NOT NULL COMMENT '头像',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '话题表';

INSERT INTO
    `topics`
VALUES
    (1, '青训营开课啦', '快到碗里来！', '2022-04-01 13:50:19');

DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `parent_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '父id',
    `content` text NOT NULL COMMENT '头像',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX parent_id (`parent_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '回帖表';

INSERT INTO
    `posts`
VALUES
    (1, 1, '举手报名！', '2022-04-01 14:50:19'),
    (2, 1, '举手报名+1', '2022-04-01 14:51:19');