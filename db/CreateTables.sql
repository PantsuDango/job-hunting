CREATE DATABASE `job-hunting`;

USE `job-hunting`;

CREATE TABLE `user` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `nick` varchar(32) NOT NULL DEFAULT '初始昵称' COMMENT '用户昵称',
   `username` varchar(32) NOT NULL COMMENT '登录账号',
   `password` varchar(32) NOT NULL COMMENT '登录密码, 存md5',
   `salt` varchar(32) NOT NULL COMMENT '盐',
   `sex` tinyint(4) NOT NULL DEFAULT 0 COMMENT '性别: 0-保密, 1-男, 2-女',
   `head_image` varchar(4096) DEFAULT NULL COMMENT '头像图片url',
   `email` varchar(32) DEFAULT NULL COMMENT '邮箱',
   `phone` varchar(32) DEFAULT NULL COMMENT '手机号',
   `birthday` varchar(32) DEFAULT NULL COMMENT '生日',
   `job` varchar(32) DEFAULT NULL COMMENT '职位',
   `Address` varchar(32) DEFAULT NULL COMMENT '家庭住址',
   `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT  '账号状态: 0-存续, 1-废除',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   `lastupdate` datetime NOT NULL COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';