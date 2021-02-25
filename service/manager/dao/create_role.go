package dao

import (
	"errors"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//CreateRole 创建角色
func (c *Conn) CreateRole(req *managerpb.Role) (*managerpb.Response, error) {
	db := c.Mysql.DB()

	sql := `INSERT IGNORE INTO manager_role (role_name) VALUES (?)`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(req.RoleName)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("该角色已创建")
	}

	return &managerpb.Response{
		Success: true,
	}, nil
}
