package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// open gzip file
	fr, err := os.Open("radius_c75a938a-9e22-40b5-98d5-873fa58aa9ec20191219154033.json.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// create gzip.Reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		log.Fatalln(err)
	}
	defer gr.Close()

	//buf := make([]byte, 1024*1024*10)
	//n, err := gr.Read(buf)

	//fw, err := os.Create(gr.Header.Name)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//_, err = fw.Write(buf[:n])
	//if err != nil {
	//	log.Fatalln(err)
	//}
	// 读取gzip对象内容
	rBuf := bufio.NewReaderSize(gr, 8192)
	for {
		data, _, eof := rBuf.ReadLine()
		if eof == io.EOF {
			break
		}
		// 以文本形式输出
		fmt.Printf("%s\n", data)
	}

}
