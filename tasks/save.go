package tasks

import (
	"sync"
)

var IpPortResultSyncMap *sync.Map

func init() {
	IpPortResultSyncMap = &sync.Map{}
}

func SaveResult(ip string, port int) {
	v, ok := IpPortResultSyncMap.Load(ip)
	if ok {
		ports, ok1 := v.([]int)
		if ok1 {
			ports = append(ports, port)
			IpPortResultSyncMap.Store(ip, ports)
		}
	} else {
		ports := make([]int, 0)
		ports = append(ports, port)
		IpPortResultSyncMap.Store(ip, ports)
	}
}
