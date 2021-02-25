package dao

import (
	accountpb "shakebook/service/account/proto/api/v1"
	"time"
)

//GetAccountList implement Conn
func (c *Conn) GetAccountList() ([]*accountpb.AccountInfo, error) {
	db := c.Mysql.DB()
	sql := `SELECT 
	id,
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
	fans_number
	FROM account`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	res := []*accountpb.AccountInfo{}
	for rows.Next() {
		item := &accountpb.AccountInfo{}
		var signupat time.Time
		var lastime time.Time
		rows.Scan(
			&item.Id,
			&item.AccountName,
			&item.AccountPhone,
			&item.AccountEmail,
			&signupat,
			&lastime,
			&item.AccountStatus,
			&item.ImageUrl,
			&item.BackgroundUrl,
			&item.AccountDesc,
			&item.ThumbsUp,
			&item.FocusNumber,
			&item.FansNumber,
		)
		item.SignupAt = signupat.Format("2006-01-02 15:04:05")
		item.LastTime = lastime.Format("2006-01-02 15:04:05")
		res = append(res, item)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
