package dao

import (
	accountpb "shakebook/service/account/proto/api/v1"
)

//GetAccount 查询用户信息
func (c *Conn) GetAccount(req *accountpb.AccountId) (*accountpb.AccountInfo, error) {
	db := c.Mysql.DB()

	sql := `SELECT 
	account_name,
	account_phone,
	account_email,
	signup_at,
	last_time,
	account_status,
	image_url,
	background_url,
	account_desc,
	thumbs_up,
	focus_number,
	fans_number,
	account_nickname
	FROM account WHERE id = ? limit 1`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	a := &accountpb.AccountInfo{}
	err = stmt.QueryRow(req.Id).Scan(
		&a.AccountName,
		&a.AccountPhone,
		&a.AccountEmail,
		&a.SignupAt,
		&a.LastTime,
		&a.AccountStatus,
		&a.ImageUrl,
		&a.BackgroundUrl,
		&a.AccountDesc,
		&a.ThumbsUp,
		&a.FocusNumber,
		&a.FansNumber,
		&a.AccountNickname,
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}
