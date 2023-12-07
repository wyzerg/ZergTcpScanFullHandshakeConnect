package tcpScanFullHandshakeConnect

import (
	"fmt"
	"testing"
)

func Test_TcpScanFullHandshakeConnect(t *testing.T) {
	result, err := TcpScanFullHandshakeConnect("127.0.0.1", "1000-8000", 5000)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", result)
}
