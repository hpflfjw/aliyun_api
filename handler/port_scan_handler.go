package handler

import (
	"flag"
	"fmt"
	"net"
)

var startPort = flag.Int("start", 1, "Start port")
var endPort = flag.Int("end", 65535, "End port")
var ipAddress = flag.String("ip", "127.0.0.1", "ipAddress")

func GetIpAddress() []string {
	// 获取本地主机的所有 IP 地址
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ipSlices := []string{}

	// 遍历所有 IP 地址
	for _, a := range addrs {
		// 将 IP 地址转换为字符串
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println(ipNet.IP.String())
				ipSlices = append(ipSlices, ipNet.IP.String())
			}
		}
	}
	return ipSlices
}

func ScanPort() {
	for i := 1; i <= 65535; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Port is closed or filtered
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
