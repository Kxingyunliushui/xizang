package main

import (
	_ "github.com/go-sql-driver/mysql"
	"xizang/src/mysql"
)



const (
	userName = "onlineinfo"
	password = "onlineinfo@swpu"
	ip       = "172.16.245.13"
	port     = "3306"
	dbName   = "srun4k"
	table    = "onlineinfo"
)


func init() {
	mysql.InitDB(userName, password, port, ip, dbName)
}

func main() {

	sql := "select user_name,ip from online_radius"
	mysql.SelectDB(sql)

	//dbinsert, _ := db.Exec("insert into zhaobsh(id,name) values('2019041901', 'zhaobsh01')")
	////执行插入的数据, db.Exec 的函数
	//fmt.Println(dbinsert);
	mysql.CloseDB()
}













//package main
//
//import (
//    "xizang/udp"
//)
//
//func main() {
//    address := "0.0.0.0:61440"
//    //for {
//    //    fileName := "/home/yxy/go/src/xizang/src/radius.json"
//    //    info, _ := json.Marshal(address)
//    //    file.Write(string(info)+"\n", fileName)
//    //}
//
//    udp.UdpServer(address)
//}