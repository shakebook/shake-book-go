package dao

import (
	accountpb "shakebook/service/account/proto/api/v1"
	managerpb "shakebook/service/manager/proto/api/v1"
	"time"
)

//GetAccountMenu implement Conn
func (c *Conn) GetAccountMenu(acc *accountpb.AccountId) ([]*managerpb.MenuInfo, error) {

	db := c.Mysql.DB()
	sql := `SELECT m.id,m.menu_name,m.menu_router,m.parent_id,m.menu_icon,m.create_at
	FROM account a
	INNER JOIN manager_account_bind_role ar
	ON ar.account_id = a.id
	INNER JOIN manager_role r
	ON r.id = ar.role_id
	INNER JOIN manager_role_bind_menu rm
	ON rm.role_id = r.id
	INNER JOIN manager_menu m
	ON m.id = rm.menu_id
	WHERE a.id=? AND m.menu_status=?
	GROUP BY m.id`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(acc.Id, managerpb.Status_NORMAL)
	if err != nil {
		return nil, err
	}

	res := []*managerpb.MenuInfo{}
	for rows.Next() {
		menu := &managerpb.MenuInfo{}
		var createtime time.Time
		rows.Scan(&menu.Id, &menu.MenuName, &menu.MenuRouter, &menu.ParentId, &menu.MenuIcon, &createtime)
		menu.CreateAt = createtime.Format("2006-01-02 15:04:05")
		res = append(res, menu)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
