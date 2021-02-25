package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
)

//AccountBindRole implement Conn
// `INSERT INTO manager_account_bind_role
// 	(account_id, role_id)
// 	VALUES (?, ?)
// 	ON DUPLICATE KEY UPDATE
// 	account_id = ?,
// 	role_id = ?`
func (c *Conn) AccountBindRole(req *managerpb.AccountBindRoleRequest) error {
	db := c.Mysql.DB()
	tx, err := db.Begin() //事务
	if err != nil {
		return err
	}

	stmt1, err := tx.Prepare(`DELETE FROM manager_account_bind_role WHERE account_id = ?`)
	defer stmt1.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err = stmt1.Exec(req.AccountId); err != nil {
		tx.Rollback()
		return err
	}

	stmt2, err := tx.Prepare(`INSERT IGNORE INTO manager_account_bind_role (account_id,role_id) VALUES (?,?)`)
	defer stmt2.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, roleid := range req.RoleIds {
		_, err := stmt2.Exec(req.AccountId, roleid)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit() //提交事务
	return nil
}
