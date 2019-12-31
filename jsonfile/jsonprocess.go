package jsonfile

import (
	"bufio"
	"encoding/json"
	"github.com/wonderivan/logger"
	"io"
	"io/ioutil"
	"os"
	"path"
	"smartvi/engine"
)

type Pgvalue struct {
	Count   int
	DbCount int
	Name    []string
}

var SrcMacmap = make(map[string]Pgvalue)

func ProcessJsonFile(pathname string) error {

	rd, err := ioutil.ReadDir(pathname)
	if nil == err {
		for _, fi := range rd {
			if fi.IsDir() {
				ProcessJsonFile(pathname + "/" + fi.Name())
			} else {
				fileSuffix := path.Ext(fi.Name())
				if ".json" == fileSuffix {
					JsonFileReadLine(pathname + "/" + fi.Name())
				}
			}
		}
	} else {
		logger.Error("ReadDir", err)
	}

	return err
}

func JsonFileReadLine(filename string) int {
	var Line_num int = 0

	fd, err := os.Open(filename)
	defer fd.Close()
	if err != nil {
		logger.Error("Open File ", filename, err)
		return -1
	}

	buff_cdr := bufio.NewReaderSize(fd, 8192)
	for {
		data, _, eof := buff_cdr.ReadLine()
		if eof == io.EOF {
			break
		}
		Line_num++
		JsonProcessJsonLine(string(data), Line_num)
	}
	return 0

}

func JsonProcessJsonLine(jsonData string, Line_num int) {
	vid := engine.Vidata{}
	err := json.Unmarshal([]byte(jsonData), &vid)
	if err != nil {
		logger.Error("L-%d,JsonDec <%s>,%s", Line_num, jsonData, err)
		return
	}
	if vid.Mac != "" {
		var tmpPgvalue Pgvalue
		tmpPgvalue.Count = SrcMacmap[vid.Mac].Count + 1
		SrcMacmap[vid.Mac] = tmpPgvalue
	}

}
