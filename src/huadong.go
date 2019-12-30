package main

//华东交通大学

import (
	"crypto/md5"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "Yplsec.com"
	ip       = "127.0.0.1"
	port     = "3336"
	dbName   = "subject_identity"
	table    = "user_data"
)

func GetMd5(src_str string) string {
	m := md5.New()
	m.Write([]byte(src_str))
	y := m.Sum(nil)
	return fmt.Sprintf("%x", y)
}

func main() {
	//sql := "select account,name,class,mobile,add_time,update_time from user_data"
	//mysql.InitDB(userName, password, port, ip, dbName)
	//mysql.SelectDB_hd(sql)
	//
	//mysql.CloseDB()
	did := "c75a938a-9e22-40b5-98d5-873fa58aa9ec" //huadong
	//did := "c082fc90-dbd3-40a1-bd81-109289986a0c"//xizang
	//did := "f8c3335f-6dbe-4324-b3b7-8e245b5931e2"//xizang
	srcip := "10.32.179.147"
	//srcip := "10.32.231.209"
	user_key := did + srcip
	key_md5 := GetMd5(user_key)
	fmt.Println(key_md5)
}
