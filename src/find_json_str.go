package main

import (
	"flag"
	"fmt"
	"xizang/jsonfile"
	"xizang/pgsql"
)

func flaginit(path *string) {
	//命令行是：test.exe -u root -p root123 -h localhost -port 8080
	//var path, filename string
	flag.StringVar(path, "path", "./", "file path")

	flag.Parse() //解析注册的flag，必须

	return
}

func main() {
	const (
		pgname = "dpi_user"
		pgpw   = "Yplsec.com"
		pghost = "10.0.3.30"
		dbport = 5432
		dbname = "campus3"
		ftable = "user_data"
	)
	var path string
	flaginit(&path)
	pgsql.PgsqlOpen(pghost, pgname, pgpw, dbname, dbport)

	jsonfile.ProcessJsonFile(path)
	for srcmac, pgvalue := range jsonfile.SrcMacmap {
		pgvalue.DbCount = pgsql.SelectSrcMac(srcmac)
		if pgvalue.DbCount != 0 {
			var nametmp []string
			name := pgsql.SelectSrcMacName(srcmac)
			for _, name := range name {
				nametmp = append(nametmp, name)
			}
			pgvalue.Name = nametmp

			jsonfile.SrcMacmap[srcmac] = pgvalue
		}
	}
	for srcmac, pgvalue := range jsonfile.SrcMacmap {
		str := fmt.Sprintf("srcmac %s count %d dbcount %d namecount %d name %v", srcmac, pgvalue.Count, pgvalue.DbCount, len(pgvalue.Name), pgvalue.Name)
		fmt.Println(str)
	}

	pgsql.PgsqlClose()
}
