package tcpScan

import (
	"context"
	"fmt"
	"github.com/wyzerg/scan/tasks"
	"github.com/wyzerg/scan/until"
)

func TCPScan(ipString, portString string, thread int) (map[string][]int, error) {
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
