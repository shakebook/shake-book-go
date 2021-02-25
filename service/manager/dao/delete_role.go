package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
)

//DeleteRole implement conn
func (c *Conn) DeleteRole(req *managerpb.RoleId) error {

	db := c.Mysql.DB()

	sql := `UPDATE manager_role SET role_status=? WHERE id = ?`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(managerpb.Status_DELETE, req.Id)
	if err != nil {
		return err
	}

	return nil
}
