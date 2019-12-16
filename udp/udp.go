package udp

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"sgygithup/goutils/datetime"
	"sgygithup/goutils/file"
	"strconv"
	"strings"
	"sync"
	"time"
)

//var MapRWMutex *sync.RWMutex

//var accountMap map[string]*radius_info
var accountMap sync.Map

// 限制goroutine数量
var limitChan = make(chan bool, 1000)

func init() {
	//accountMap = make(map[string]*radius_info)
	Timer()
}

// UDP goroutine 实现并发读取UDP数据
func udpProcess(conn *net.UDPConn) {

	// 最大读取数据大小
	data := make([]byte, 246)
	_, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed read udp msg, error: " + err.Error())
	}

	fmt.Printf("data:[% x]\n", data)
	reaDdata(data)
	udpWrite(conn, remoteAddr)
	<-limitChan
}

func udpWrite(conn *net.UDPConn, addr *net.UDPAddr) {
	data := []byte{0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00}
	conn.WriteToUDP(data, addr)
}

type radius_info struct {
	User_name   string `json:"user_name, string"`
	Account     string `json:"account, string"`
	Userip      string `json:"userip, string"`
	Mac_address string `json:"mac_address, string"`
	Logintime   string `json:"logintime, string"`
}

func ParseRtuDevid(input []byte) (output int64) {
	for i := len(input) - 1; i >= 0; i-- {
		output = (output << 8) + int64(uint32(input[i]))
	}
	return
}

func Timer() {
	ticker := time.NewTicker(time.Second * 60 * 5) // 5 分钟
	go func() {
		for _ = range ticker.C {
			WriteJson()
		}

	}()
}

func reaDdata(data []byte) {
	data_info := &radius_info{}
	account := strings.Split(string(data[32:64]), "@")[0]
	account = strings.Split(account, "\u0000")[0]
	data_info.Account = account
	data_info.Mac_address = toMac(data[231:237])
	data_info.Userip = toIp(data[227:231])
	//data_info.User_name = string(data[64:96])
	data_info.User_name = ""
	timestamp := datetime.DateTotime(string(data[12:31]))
	logtime := strconv.FormatInt(timestamp, 10)
	data_info.Logintime = logtime
	cmd := int(data[226])
	if cmd == 1 || cmd == 3 { // 1 上线 3 在线
		accountMap.Store(account, data_info)
	} else if cmd == 2 { // 2 下线 暂不处理
		accountMap.Delete(account)
	}

	//fileName := "/home/dpiuser/radius/radius.json"
	//info, _ := json.Marshal(data_info)
	//file.Write(string(info)+"\n", fileName)
	//fmt.Println("命令代码", ParseRtuDevid(data[:4]))
	//fmt.Printf("命令代码:[% x]\n", data[:4])
	//fmt.Println("子命令代码",ParseRtuDevid(data[4:8]))
	//fmt.Println("发送的数据长度",ParseRtuDevid(data[8:12]))
	//fmt.Println("事件发生时间",string(data[12:31]))
	//fmt.Println("账号",strings.Split(string(data[32:64]), "@")[0])
	//fmt.Println("真实姓名",string(data[64:96]))
	//fmt.Println("证件类型",string(data[96:98]))
	//fmt.Println("证件号码",string(data[98:118]))
	//fmt.Println("居住、单位地址",string(data[118:150]))
	//fmt.Println("国籍",string(data[150:162]))
	//fmt.Println("备注",string(data[162:218]))
	//fmt.Println("计费组ID",string(data[218:220]))
	//fmt.Println("带宽组ID",string(data[220:222]))
	//fmt.Println("行政组ID",string(data[222:224]))
	//fmt.Println("协议主版本号",string(data[224:226]))
	//fmt.Println("协议次版本号",string(data[226:228]))
	//fmt.Println("校验和",string(data[228:232]))
	//fmt.Println("事件类型",string(data[226:227]))
	//fmt.Printf("事件类型:[% x]\n", data[226:227])
	//fmt.Println("登录IP地址",toIp(data[227:231]))
	//fmt.Printf("登录IP地址:[% x]\n", data[227:231])
	//fmt.Println("登录MAC地址",toMac(data[231:237]))
	//fmt.Printf("登录MAC地址:[% x]\n", data[231:236])
	//fmt.Println("备用",string(data[243:246]))

	//outstring := fmt.Sprintf("命令代码:[% x]\n", data[:4]) +
	//fmt.Sprintf("子命令代码",ParseRtuDevid(data[4:8]))  +
	//fmt.Sprintf("发送的数据长度",ParseRtuDevid(data[8:12])) +
	//fmt.Sprintf("事件发生时间",string(data[12:31])) +
	//fmt.Sprintf("账号",strings.Split(string(data[32:64]), "@")[0]) +
	//fmt.Sprintf("真实姓名",string(data[64:96])) +
	//fmt.Sprintf("证件类型",string(data[96:98])) +
	//fmt.Sprintf("证件号码",string(data[98:118])) +
	//fmt.Sprintf("居住、单位地址",string(data[118:150])) +
	//fmt.Sprintf("国籍",string(data[150:162])) +
	//fmt.Sprintf("备注",string(data[162:218])) +
	//fmt.Sprintf("计费组ID",string(data[218:220])) +
	//fmt.Sprintf("带宽组ID",string(data[220:222])) +
	//fmt.Sprintf("行政组ID",string(data[222:224])) +
	//fmt.Sprintf("协议主版本号",string(data[224:226])) +
	//fmt.Sprintf("协议次版本号",string(data[226:228])) +
	//fmt.Sprintf("校验和",string(data[228:232])) +
	//fmt.Sprintf("事件类型",string(data[226:227])) +
	//fmt.Sprintf("事件类型:[% x]\n", data[226:227])+
	//	fmt.Sprintf("登录IP地址",toIp(data[227:231])) +
	//fmt.Sprintf("登录IP地址:[% x]\n", data[227:231])+
	//	fmt.Sprintf("登录MAC地址",toMac(data[231:237]))+
	//	fmt.Sprintf("登录MAC地址:[% x]\n", data[231:236])+
	//	fmt.Sprintf("备用",string(data[243:246]))
	//fileName2 := "/home/dpiuser/radius/radius.json"
	//info2, _ := json.Marshal(outstring)
	//file.Write(string(info2)+"\n", fileName2)

}

func WriteJson() {
	//uuid := "c75a938a-9e22-40b5-98d5-873fa58aa9ec_" //
	uuid := "c082fc90-dbd3-40a1-bd81-109289986a0c_" //西藏民族大学
	//route := "/home/radius/"
	route := "/home/dpiuser/radius/"
	name := "radius_" + uuid + Gettime() + ".json"
	filename := route + name
	f := file.CreateFile(filename)
	defer f.Close()
	fc := func(key, value interface{}) bool {
		//fmt.Printf("Range: k, v = %v, %v\n", key, value)
		info, _ := json.Marshal(value)
		file.WriteData(string(info)+"\n", f)
		return true
	}
	accountMap.Range(fc)
	//上传
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

func Gettime() string {
	t := time.Now().Unix()
	tmt := time.Unix(t, 0).Format("20060102150405")
	return tmt
}

func toIp(data []byte) (str string) {
	return fmt.Sprintf("%d.%d.%d.%d", int(data[0]), int(data[1]), int(data[2]), int(data[3]))
}

func toMac(data []byte) (str string) {
	str = ""
	for _, value := range data {
		if value < 10 {
			str += "0" + strconv.FormatInt(int64(value), 16)
		} else {
			str += strconv.FormatInt(int64(value), 16)
		}
	}
	return strings.ToUpper(str)
}

func UdpServer(address string) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", udpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("read from connect failed, err:" + err.Error())
		os.Exit(1)
	}
	fmt.Println("Start collecting ......")
	for {
		limitChan <- true
		go udpProcess(conn)
	}
}
