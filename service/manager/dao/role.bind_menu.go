package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
)

//RoleBindMenu implement Conn
func (c *Conn) RoleBindMenu(req *managerpb.RoleBindMenuRequest) error {
	db := c.Mysql.DB()
	tx, err := db.Begin() //事务
	if err != nil {
		return err
	}

	stmt1, err := tx.Prepare(`DELETE FROM manager_role_bind_menu WHERE role_id = ?`)
	defer stmt1.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err = stmt1.Exec(req.RoleId); err != nil {
		tx.Rollback()
		return err
	}

	stmt2, err := tx.Prepare(`INSERT IGNORE INTO manager_role_bind_menu (role_id,menu_id) VALUES (?,?)`)
	defer stmt2.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, menuid := range req.MenuIds {
		_, err := stmt2.Exec(req.RoleId, menuid)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit() //提交事务
	return nil
}
