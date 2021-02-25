


CREATE TABLE manager_account_bind_role
(
   account_id bigint(20) NOT NULL COMMENT '账号id',
   role_id int NOT NULL COMMENT '角色id',
   PRIMARY KEY(account_id,role_id),
   KEY index_account_id (account_id) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

SELECT a.id, r.id
FROM account a
INNER JOIN manager_account_bind_role ar
ON ar.account_id = a.id
INNER JOIN manager_role r
ON r.id = ar.role_id;

查询账号菜单:
SELECT m.id,m.menu_name,m.menu_router,m.parent_id,m.menu_icon,m.create_at,m.menu_status
FROM account a
INNER JOIN manager_account_bind_role ar
ON ar.account_id = a.id
INNER JOIN manager_role r
ON r.id = ar.role_id
INNER JOIN manager_role_bind_menu rm
ON rm.role_id = r.id
INNER JOIN manager_menu m
ON m.id = rm.menu_id
WHERE a.id=1
GROUP BY m.id;




