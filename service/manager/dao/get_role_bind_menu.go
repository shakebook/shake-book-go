package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
)

//GetRoleBindMenu implement conn
func (c *Conn) GetRoleBindMenu(req *managerpb.RoleId) ([]int32, error) {
	db := c.Mysql.DB()
	sql := `SELECT menu_id FROM manager_role_bind_menu where role_id = ?`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(req.Id)
	if err != nil {
		return nil, err
	}

	var res []int32
	for rows.Next() {
		var menuid int32
		rows.Scan(&menuid)
		res = append(res, menuid)
	}

	return res, nil
}
