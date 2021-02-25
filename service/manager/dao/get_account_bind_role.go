package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
)

//GetAccountBindRole implement conn
func (c *Conn) GetAccountBindRole(req *managerpb.AccountId) ([]int32, error) {
	db := c.Mysql.DB()
	sql := `SELECT role_id FROM manager_account_bind_role WHERE account_id = ?`
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
		var roleid int32
		rows.Scan(&roleid)
		res = append(res, roleid)
	}

	return res, nil
}
