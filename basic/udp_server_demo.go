package basic

import (
	"fmt"
	"net"
	"os"
)

// 限制goroutine数量
var limitChan = make(chan bool, 1000)

// UDP goroutine 实现并发读取UDP数据
func UdpProcess(conn *net.UDPConn) {

	// 最大读取数据大小
	data := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed read udp msg, error: " + err.Error())
	}
	str := string(data[:n])
	fmt.Println("receive from client, data:" + str)
	<-limitChan
}

func UdpServer(address string) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", udpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("read from connect failed, err:" + err.Error())
		os.Exit(1)
	}

	for {
		limitChan <- true
		go UdpProcess(conn)
	}
}
