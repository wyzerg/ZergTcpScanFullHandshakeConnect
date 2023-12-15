package main

import (
	"context"
	"fmt"
	"github.com/wyzerg/tcpScan/tasks"
	"github.com/wyzerg/tcpScan/until"
)

func TCPScan(ipString, portString string, thread int) (map[string][]int, error) {
	var ctx, cancel = context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	ips, err := until.ParseIps(ipString)
	if err != nil {
		return nil, fmt.Errorf("TcpScan err:%v\n", err)
	}
	ports, err := until.ParsePorts(portString)
	if err != nil {
		return nil, fmt.Errorf("TcpScan err:%v\n", err)
	}
	result := tasks.Run(ips, ports, thread, ctx)
	return result, nil
}

func main() {
	result, err := TCPScan("127.0.0.1", "1000-10000", 5000)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", result)
}
