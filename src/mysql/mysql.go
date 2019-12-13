package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"sgygithup/goutils/file"
	"strings"
)

//Db数据库连接池
var DB *sql.DB

type sql_buf struct {
	User_name string `json:"user_name, string"`
	Ip        string `json:"ip, string"`
}


//注意方法名大写，就是public
func InitDB(userName, password, port, ip, dbName string) {

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func SelectDB(sql string) {
	rows, _ := DB.Query(sql)
	//rows 查询 表里面所有的数据 结果应该是一个数组 方式db.Query
	user_name := ""
	ip := ""
	data_info := &sql_buf{}
	fileName := "/home/dpiuser/sqltest/user_ip.txt"

	for rows.Next() {
		rows.Scan(&user_name, &ip)
		data_info.User_name = user_name
		data_info.Ip = ip
		info, _ := json.Marshal(data_info)
		file.Write(string(info)+"\n", fileName)
		//fmt.Println(user_name, ip);
	}
	////遍历数组里面的内容. 并且打印出来.  Scan 和 Next 的函数

}

func CloseDB()  {
	if DB != nil {
		DB.Close()
	}
}