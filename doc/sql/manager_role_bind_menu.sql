
CREATE TABLE manager_role_bind_menu
(
   role_id int NOT NULL COMMENT '角色id',
   menu_id int NOT NULL COMMENT '菜单id',
   PRIMARY KEY(role_id,menu_id),
   KEY index_menu_id (role_id) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
