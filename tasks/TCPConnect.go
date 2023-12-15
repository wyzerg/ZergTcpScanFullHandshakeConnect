package tasks

import (
	"fmt"
	"net"
	"time"
)

func TCPConnect(ip string, port int) error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), 2*time.Second)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return err
}
