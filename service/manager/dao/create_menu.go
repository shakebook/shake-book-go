package dao

import (
	"errors"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//CreateMenu implement Conn
func (c *Conn) CreateMenu(req *managerpb.CreateMenuRequest) error {
	db := c.Mysql.DB()

	sql := `INSERT IGNORE INTO manager_menu 
					(
					menu_name,
					menu_router,
					parent_id,
					menu_icon 
					)
					VALUES (?,?,?,?)`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(
		req.MenuName,
		req.MenuRouter,
		req.ParentId,
		req.MenuIcon,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("创建失败")
	}

	return nil
}
