package dao

import (
	"context"
	"errors"
	accountpb "shakebook/service/account/proto/api/v1"

	"golang.org/x/crypto/bcrypt"
)

//SignIn dao
func (c *Conn) SignIn(ctx context.Context, req *accountpb.SignInRequest) (*accountpb.AccountId, error) {
	db := c.Mysql.DB()
	stmt, err := db.Prepare(`SELECT id, account_password FROM account WHERE account_name = ? limit 1`)
	if err != nil {
		return nil, err
	}
	var id int64
	var pwd string
	err = stmt.QueryRow(req.AccountName).Scan(&id, &pwd)
	if err != nil {
		return nil, errors.New("账号不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(req.AccountPassword))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	return &accountpb.AccountId{
		Id: id,
	}, nil

}
