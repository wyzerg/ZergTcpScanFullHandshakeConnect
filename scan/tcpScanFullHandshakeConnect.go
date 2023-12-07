package scan

import (
	"context"
	"fmt"
	"tcpScanFullHandshakeConnect/tasks"
	"tcpScanFullHandshakeConnect/until"
)

func TcpScan(ipString, portString string, thread int) (map[string][]int, error) {
	var ctx, cancel = context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	ips, err := until.ParseIps(ipString)
	if err != nil {
		return nil, fmt.Errorf("TcpScanFullHandshakeConnect err:%v\n", err)
	}
	ports, err := until.ParsePorts(portString)
	if err != nil {
		return nil, fmt.Errorf("TcpScanFullHandshakeConnect err:%v\n", err)
	}
	result := tasks.Run(ips, ports, thread, ctx)
	return result, nil
}
