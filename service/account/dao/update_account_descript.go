package dao

import (
	accountpb "shakebook/service/account/proto/api/v1"
)

//UpdateAccountDescript 更新用户信息
//account_desc = case when ? is not null and length(?) > 0 then ? else account_desc end
func (c *Conn) UpdateAccountDescript(id int64, req *accountpb.UpdateAccountDescriptRequest) error {
	db := c.Mysql.DB()
	sql := `UPDATE account SET 
	account_nickname=?,
	account_name = ?,
	account_desc = ?
	WHERE id = ?`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		req.AccountNickname,
		req.AccountName,
		req.AccountDesc,
		id,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
