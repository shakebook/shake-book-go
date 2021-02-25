
CREATE TABLE manager_menu
(
  id int NOT NULL AUTO_INCREMENT COMMENT '主键',
	menu_name char(64) NOT NULL COMMENT '菜单名',
  menu_router varchar(200) NOT NULL COMMENT '菜单路由',
	parent_id int DEFAULT 0 COMMENT '上级菜单id',
	menu_icon varchar(200) DEFAULT '' COMMENT '菜单图标',
	create_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
	menu_status int(1) DEFAULT 0 COMMENT '状态0:NORMAL,1:DELETE',
	PRIMARY KEY(id),
	KEY index_menu_status(menu_status)
) ENGINE = InnoDB CHARSET = utf8;
ALTER TABLE menu ADD KEY index_status (status) USING BTREE;


