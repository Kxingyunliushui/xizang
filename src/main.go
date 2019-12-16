//package main
//
//import (
//	_ "github.com/go-sql-driver/mysql"
//	"xizang/mysql"
//)
//
//
//
//const (
//	userName = "onlineinfo"
//	password = "onlineinfo@swpu"
//	ip       = "172.16.245.13"
//	port     = "3306"
//	dbName   = "srun4k"
//	table    = "onlineinfo"
//)
//
//
//func main() {
//
//	sql := "select user_name,ip,user_mac,add_time from online_radius"
//
//	mysql.InitDB(userName, password, port, ip, dbName)
//	mysql.Timer(sql)
//
//	select {} // 阻塞
//
//
//
//	//sql := "select * from online_radius"
//	//var sql string
//	//for num, arg := range os.Args {
//	//	if num == 0 {
//	//		continue
//	//	}
//	//	sql += arg
//	//	if num < len(os.Args)-1 {
//	//		sql += " "
//	//	}
//	//}
//
//	//fmt.Println(sql)
//
//	//
//	//results := mysql.SelectDBprint(sql)
//	//for _, result := range results {
//	//	for name, value := range result {
//	//		fmt.Printf("{%s\t%s}\n", name, value)
//	//	}
//	//	fmt.Printf("\n")
//	//	//break
//	//
//	//}
//
//	//dbinsert, _ := db.Exec("insert into zhaobsh(id,name) values('2019041901', 'zhaobsh01')")
//	////执行插入的数据, db.Exec 的函数
//	//fmt.Println(dbinsert);
//	mysql.CloseDB()
//}

//uuid := "c082fc90-dbd3-40a1-bd81-109289986a0c_" //西藏民族大学
//uuid := "0c2aad63-24ee-4d08-a8af-cf3d331e9849_" //西南石油大学

package main

import (
	"xizang/udp"
)

func main() {
	address := "0.0.0.0:61440"
	//for {
	//    fileName := "/home/yxy/go/src/xizang/src/radius.json"
	//    info, _ := json.Marshal(address)
	//    file.Write(string(info)+"\n", fileName)
	//}

	udp.UdpServer(address)
}
