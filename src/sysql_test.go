//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"strings"
//)
//
////Db数据库连接池
//var DB *sql.DB
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
//type sql_buf struct {
//	User_name string `json:"user_name, string"`
//	Ip        string `json:"ip, string"`
//}
//
//func init() {
//	InitDB()
//}
//
//func main() {
//
//	rows, _ := DB.Query("select user_name,ip from online_radius")
//	//rows 查询 表里面所有的数据 结果应该是一个数组 方式db.Query
//	user_name := ""
//	ip := ""
//	data_info := &sql_buf{}
//	fileName := "/home/dpiuser/sqltest/user_ip.txt"
//
//	for rows.Next() {
//		rows.Scan(&user_name, &ip)
//		data_info.User_name = user_name
//		data_info.Ip = ip
//		//fmt.Println(user_name, ip);
//	}
//	////遍历数组里面的内容. 并且打印出来.  Scan 和 Next 的函数
//
//	info, _ := json.Marshal(data_info)
//	file.Write(string(info)+"\n", fileName)
//	//dbinsert, _ := db.Exec("insert into zhaobsh(id,name) values('2019041901', 'zhaobsh01')")
//	////执行插入的数据, db.Exec 的函数
//	//fmt.Println(dbinsert);
//	DB.Close()
//}
//
////注意方法名大写，就是public
//func InitDB() {
//	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
//	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
//
//	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
//	DB, _ = sql.Open("mysql", path)
//	//设置数据库最大连接数
//	DB.SetConnMaxLifetime(100)
//	//设置上数据库最大闲置连接数
//	DB.SetMaxIdleConns(10)
//	//验证连接
//	if err := DB.Ping(); err != nil {
//		fmt.Println("opon database fail")
//		return
//	}
//	fmt.Println("connnect success")
//}
