package pgsql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Pgdb *sql.DB

func PgsqlOpen(host, user, password, dbname string, port int) {
	var err error
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Pgdb, err = sql.Open("postgres", pgInfo)
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;

	//还可以是这种方式打开
	//db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	checkErr(err)
}

func PgsqlClose() {
	Pgdb.Close()
}

func SelectSrcMac(srcmac string) int {
	var count int = 0
	//查询数据
	sqlStatement := `SELECT COUNT(mac) FROM subject_identity.user_info WHERE did='a2b3d758-8681-4659-a9e9-2f25c33c0b30' and mac = '` + srcmac + `'`
	rows, err := Pgdb.Query(sqlStatement)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func SelectSrcMacName(srcmac string) []string {
	var namesplit []string
	var name string
	var namemap = make(map[string]int)

	//查询数据
	sqlStatement := `SELECT name FROM subject_identity.user_info WHERE did='a2b3d758-8681-4659-a9e9-2f25c33c0b30' and mac = '` + srcmac + `'`
	rows, err := Pgdb.Query(sqlStatement)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&name)
		if name != "" {
			namemap[name] = 1
		}
		checkErr(err)
	}

	for name, _ := range namemap {
		namesplit = append(namesplit, name)
	}
	return namesplit
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
