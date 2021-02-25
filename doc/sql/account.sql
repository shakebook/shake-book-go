CREATE TABLE account
(
   id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
   account_name varchar(20) NOT NULL COMMENT '用户名',
   account_password varchar(64) NOT NULL COMMENT '用户密码',
   account_phone varchar(20) DEFAULT '' COMMENT '手机号码',
   account_email varchar(64) NOT NULL COMMENT '邮箱',
   signup_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
   last_time datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间',
   account_status int DEFAULT 0 COMMENT '账户状态(NORMAL/DISABLE/LOCK/DELETE)',
   image_url varchar(2083) DEFAULT '' COMMENT '头像',
   background_url varchar(2083) DEFAULT '' COMMENT '头像',
   account_desc varchar(100) DEFAULT '' COMMENT '简介',
   thumbs_up bigint(20) DEFAULT 0 COMMENT '点赞数',
   focus_number bigint(20) DEFAULT 0 COMMENT '喜欢数',
   fans_number bigint(20) DEFAULT 0 COMMENT '粉丝数',
   PRIMARY KEY(id),
   UNIQUE KEY index_account_name (account_name) USING BTREE,
   UNIQUE KEY index_account_email (account_email) USING BTREE,
   UNIQUE KEY index_account_phone (account_phone) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

drop index index_account_phone on tableName;
ALTER TABLE account ADD COLUMN account_nickname VARCHAR(10) AFTER fans_number;