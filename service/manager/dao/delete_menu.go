package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
)

//DeleteMenu implement Conn
func (c *Conn) DeleteMenu(req *managerpb.MenuId) error {
	db := c.Mysql.DB()
	sql := `UPDATE manager_menu SET menu_status=? WHERE id = ?`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		managerpb.Status_DELETE,
		req.Id,
	)
	if err != nil {
		return err
	}

	return nil
}
