package main

//华东交通大学

import (
	_ "github.com/go-sql-driver/mysql"
	"xizang/mysql"
)

const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "subject_identity"
	table    = "user_data"
)

func main() {
	sql := "select name,class,mobile,add_time,update_time from user_data"
	mysql.InitDB(userName, password, port, ip, dbName)
	mysql.SelectDB_hd(sql)

	mysql.CloseDB()
}
