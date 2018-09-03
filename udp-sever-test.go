package main

import (
	"fmt"
	"net"
)

func main() {
	var receivedCounter int64 = 0
	addrStr := net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 1110}
	socket, err := net.ListenUDP("udp4", &addrStr)
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer socket.Close()
	for {
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		fmt.Println("received msg!!!")
		if err != nil {
			fmt.Println("read data err,", err)
			continue
		}
		//data[read]='\0'
		fmt.Println(read, remoteAddr)
		receivedCounter++
		fmt.Printf("receved conuter=%d\r\n", receivedCounter)
		var pr []uint8 = data[:20]
		if read > 20 {
			fmt.Printf("received:%s...\r\n", pr)
		} else {
			fmt.Printf("received:%s\r\n", data)
		}

		_, err = socket.WriteToUDP(data, remoteAddr)
		if err != nil {
			fmt.Println("write udp  error")
			return
		}
	}
}
