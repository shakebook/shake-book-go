package dao

import managerpb "shakebook/service/manager/proto/api/v1"

//UpdateRole implement Conn
//如果数据表为not null 使用以下方式更新
func (c *Conn) UpdateRole(req *managerpb.Role) error {
	db := c.Mysql.DB()

	sql := `UPDATE manager_role SET 
	role_name = case when ? is not null and length(?) > 0 then ? else role_name end,
	role_status = case when ? is not null and length(?) > 0 then ? else role_status end
	WHERE id = ?`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		req.RoleName, req.RoleName, req.RoleName,
		req.RoleStatus, req.RoleStatus, req.RoleStatus,
		req.Id,
	)
	if err != nil {
		return err
	}

	return nil
}
