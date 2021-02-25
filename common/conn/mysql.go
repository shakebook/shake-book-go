package conn

import (
	"database/sql"
	"log"
	"strings"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Mysql 实现者
type Mysql struct {
	UserName     string
	Password     string
	Addr         string
	DatabaseName string
}

var db *sql.DB

//DB implement Mysql
func (c *Mysql) DB() *sql.DB {
	if db != nil {
		return db
	}

	url := connectURL(c)
	db, _ = sql.Open("mysql", url)
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to mysql, err:%v\n", err.Error())
	}
	log.Println("Connect to mysql successfuly")
	return db
}

func connectURL(c *Mysql) string {
	var b strings.Builder
	b.WriteString(c.UserName)
	b.WriteString(":")
	b.WriteString(c.Password)
	b.WriteString("@tcp(")
	b.WriteString(c.Addr)
	b.WriteString(")/")
	b.WriteString(c.DatabaseName)
	b.WriteString("?charset=utf8&parseTime=True")
	return b.String()
}
