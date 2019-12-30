package main

//根据redis文件读取的信息更新user_info表
import (
	"bufio"
	"database/sql"
	"encoding/json"
	"github.com/wonderivan/logger"
	"io"
	"os"
	"smartvi/conf"
	"smartvi/db"
)

type file_buf_hd struct {
	Account     string `json:"account, string"`
	User_Name   string `json:"user_name, string"`
	Class       string `json:"class, string"`
	Mobile      string `json:"mobile, string"`
	Add_time    string `json:"add_time, string"`
	Update_time string `json:"update_time, string"`
}

const (
	sql_name = "dpi_user"
	sql_pw   = "Yplsec.com"
	//sql_host = "127.0.0.1"
	sql_host   = "10.0.3.30"
	sql_port   = 5432
	sql_dbname = "campus3"
	sql_table  = "user_data"
)

var file_map = make(map[string]file_buf_hd)
var update_buf []db.Pg_user_query
var sql_user_hd *sql.DB
var find_user_did = "c75a938a-9e22-40b5-98d5-873fa58aa9ec" //"华东交通大学"
var Sql_transaction_hd *sql.Tx

func sqlinit() int {
	sql_user_hd = conf.Pgsql_connect(sql_host, sql_port, sql_name, sql_pw, sql_dbname)
	if nil == sql_user_hd {
		logger.Error("SQL Connect Failed:", sql_host, sql_port, sql_name, sql_dbname)
		return -1
	}
	Sql_transaction_hd, _ = sql_user_hd.Begin() // 开启事务
	logger.Info("PG sql_vi_user_hd success", sql_host, sql_port, sql_name, sql_dbname)
	return 0
}

func readfile(filename string) int {
	var Line_num int
	fd, err := os.Open(filename)
	defer fd.Close()
	if err != nil {
		logger.Error("Open File ", filename, err)
		return -1
	}

	filebuff := bufio.NewReaderSize(fd, 8192)
	for {
		jsonData, _, eof := filebuff.ReadLine()
		if eof == io.EOF {
			break
		}
		Line_num++
		vid := file_buf_hd{}
		err := json.Unmarshal([]byte(jsonData), &vid)
		if err != nil {
			logger.Error("L-%d,JsonDec <%s>,%s", Line_num, jsonData, err)
			return -1
		}
		var file_buf file_buf_hd
		file_buf.Update_time = vid.Update_time
		file_buf.Account = vid.Account
		file_buf.Add_time = vid.Add_time
		file_buf.Mobile = vid.Mobile
		file_buf.Class = vid.Class
		file_buf.User_Name = vid.User_Name
		file_map[vid.Account] = file_buf

	}
	logger.Info("read file %s ok", filename)
	return 0
}

func main() {
	if -1 == sqlinit() {
		return
	}
	readfile("radius_c75a938a-9e22-40b5-98d5-873fa58aa9ec_20191219154033.json")

	userinfos := db.Vi_user_pg_null_select(Sql_transaction_hd, "did", "account", find_user_did)

	for _, userinfo := range userinfos {
		if userinfo.Phone == "" || userinfo.Name == "" {
			var tmp_update_buf db.Pg_user_query
			tmp_update_buf = userinfo
			tmp_update_buf.Name = file_map[userinfo.Account].User_Name
			tmp_update_buf.Phone = file_map[userinfo.Account].Mobile
			update_buf = append(update_buf, tmp_update_buf)
		}
	}

	for _, update := range update_buf {
		db.Vi_user_pg_update_trueId(Sql_transaction_hd, &update)
		logger.Info("update date:", update)
	}
	logger.Info("Update table user_info count is ", len(update_buf))
	err := Sql_transaction_hd.Commit() // 提交
	if err != nil {
		logger.Error("Cancel the operation", err)
		Sql_transaction_hd.Rollback() // 回滚
	}
	defer sql_user_hd.Close()
}
