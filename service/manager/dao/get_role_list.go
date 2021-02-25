package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
	"time"
)

//GetRoleList implement conn
func (c *Conn) GetRoleList() ([]*managerpb.Role, error) {

	db := c.Mysql.DB()

	sql := `SELECT id,role_name,create_at FROM manager_role where role_status = 0`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	res := []*managerpb.Role{}
	for rows.Next() {
		item := &managerpb.Role{}
		var createtime time.Time
		rows.Scan(&item.Id, &item.RoleName, &createtime)
		item.CreateAt = createtime.Format("2006-01-02 15:04:05")
		res = append(res, item)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
