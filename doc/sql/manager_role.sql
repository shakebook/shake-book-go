CREATE TABLE manager_role
(
   id int NOT NULL AUTO_INCREMENT COMMENT '主键',
   role_name varchar(64) NOT NULL DEFAULT '' COMMENT '角色名',
   create_at datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
   role_status int(1) NOT NULL DEFAULT 0 COMMENT '角色状态0:NORMAL,DELETE',
   PRIMARY KEY(id),
   KEY index_role_status (role_status)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;