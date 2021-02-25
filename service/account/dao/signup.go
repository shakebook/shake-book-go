package dao

import (
	"errors"
	"shakebook/common/tools"
	accountpb "shakebook/service/account/proto/api/v1"
)

//SignUp 注册
func (c *Conn) SignUp(a *accountpb.SignUpRequest) error {
	db := c.Mysql.DB()
	stmt, err := db.Prepare(`INSERT IGNORE INTO account (account_name, account_email, account_password,account_nickname) VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	nickname, err := tools.GenerateRandCode(8, "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		return errors.New("生成随机名称失败")
	}
	res, err := stmt.Exec(a.AccountName, a.AccountEmail, a.AccountPassword, nickname)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("该账号已注册")
	}

	return nil
}
