package tcpScan

import (
	"fmt"
	"testing"
)

func Test_TcpConnect(t *testing.T) {
	result, err := TCPScan("127.0.0.1", "1000-10000", 5000)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", result)
}
