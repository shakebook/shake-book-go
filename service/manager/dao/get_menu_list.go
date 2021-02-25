package dao

import (
	managerpb "shakebook/service/manager/proto/api/v1"
	"time"
)

//GetMenuList implement conn
func (c *Conn) GetMenuList() ([]*managerpb.MenuInfo, error) {

	db := c.Mysql.DB()
	sql := `SELECT id,menu_name,menu_router,parent_id,menu_icon,create_at FROM manager_menu where menu_status = 0`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	res := []*managerpb.MenuInfo{}
	for rows.Next() {
		item := &managerpb.MenuInfo{}
		var createtime time.Time
		rows.Scan(&item.Id, &item.MenuName, &item.MenuRouter, &item.ParentId, &item.MenuIcon, &createtime)
		item.CreateAt = createtime.Format("2006-01-02 15:04:05")
		res = append(res, item)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
