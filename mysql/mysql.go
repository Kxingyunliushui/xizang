package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"time"
	"xizang/myutil"

	"sgygithup/goutils/file"
	"strings"
)

//Db数据库连接池
var DB *sql.DB

type sql_buf struct {
	User_name   string `json:"user_name, string"`
	Account     string `json:"account, string"`
	Userip      string `json:"userip, string"`
	Mac_address string `json:"mac_address, string"`
	Logintime   string `json:"logintime, string"`
}

func Gettime() string {
	t := time.Now().Unix()
	tmt := time.Unix(t, 0).Format("20060102150405")
	return tmt
}

func Timer(sql string) {
	ticker := time.NewTicker(time.Second * 60 * 5) // 5 分钟
	go func() {
		for _ = range ticker.C {
			SelectDB(sql)
		}

	}()
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
	var account, userip, mac, logintime string
	data_info := &sql_buf{}

	//fileName := "/home/dpiuser/sqltest/user_ip.txt"
	uuid := "0c2aad63-24ee-4d08-a8af-cf3d331e9849_" //西南石油大学
	//route := "/home/radius/"
	route := "/home/dpiuser/radius/"
	name := "radius_" + uuid + Gettime() + ".json"
	filename := route + name
	f := file.CreateFile(filename)

	for rows.Next() {
		rows.Scan(&account, &userip, &mac, &logintime)
		data_info.User_name = ""
		data_info.Account = account
		data_info.Userip = userip
		data_info.Mac_address = myutil.Utils_mac(mac)
		data_info.Logintime = myutil.Util_time(logintime)
		info, _ := json.Marshal(data_info)
		file.Write(string(info)+"\n", filename)
		//fmt.Println(user_name, ip);
	}
	f.Close()
	////遍历数组里面的内容. 并且打印出来.  Scan 和 Next 的函数

	//上传文件并删除本地
	url := "http://data.campus.yplsec.cn/dpilog/"
	param := make(map[string]string)
	param["data_type"] = "radius"
	param["version"] = "1.0"
	nameField := "file" // key 值
	fileName := filename
	files, _ := os.Open(fileName)
	defer files.Close()
	data, err := file.UploadFile(url, param, nameField, name, files)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("data", string(data), name)
	cmd := exec.Command("/bin/bash", "-c", "rm "+filename)
	fmt.Println("cmd", cmd)
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println("cmd rm err", err)
	}
	resp := string(bytes)
	fmt.Println("resp", resp)

}

type s1 map[string]string //make
type s2 []s1

func SelectDBprint(sql string) s2 {
	rows, _ := DB.Query(sql)
	//rows 查询 表里面所有的数据 结果应该是一个数组 方式db.Query

	if rows == nil {
		return nil
	}

	cols, err := rows.Columns()

	rawResult := make([][]byte, len(cols))
	result := s2{}
	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	//rowsNum:=0;
	for rows.Next() {
		err = rows.Scan(dest...)

		sresult := make(s1, len(cols))

		for i, raw := range rawResult {
			if raw == nil {
				sresult[cols[i]] = ""
			} else {
				sresult[cols[i]] = string(raw)
			}
		}
		result = append(result, sresult)

	}

	_ = err
	return result
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
