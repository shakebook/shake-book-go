package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
)

//UpdateMenu implement Conn
func (c *Conn) UpdateMenu(req *managerpb.UpdateMenuRequest) error {
	db := c.Mysql.DB()
	sql := `UPDATE manager_menu SET menu_name=?,menu_status=?,menu_icon=?,menu_router=?,parent_id=? WHERE id = ?`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		req.MenuName,
		req.MenuStatus,
		req.MenuIcon,
		req.MenuRouter,
		req.ParentId,
		req.Id,
	)
	if err != nil {
		return err
	}

	return nil
}
