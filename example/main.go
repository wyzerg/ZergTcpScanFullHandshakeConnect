package main

import (
	"fmt"
	"tcpScanFullHandshakeConnect"
)

func main() {
	result, err := tcpScanFullHandshakeConnect.TcpScanFullHandshakeConnect("127.0.0.1", "1000-8000", 5000)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", result)
}
