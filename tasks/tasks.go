package tasks

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type IpPort struct {
	Ip   string
	Port int
}

func Run(ips []net.IP, ports []int, thread int, ctx context.Context) map[string][]int {
	var WG sync.WaitGroup
	ipPortsCh := make(chan IpPort, 100)
	threadCh := make(chan bool, thread)
	returnChan := make(chan bool, 1)
	ResultsCh := make(chan IpPort)

	go func(ctx context.Context) {
		for {
			select {
			case <-returnChan:
				return
			case ipPort := <-ipPortsCh:
				threadCh <- true
				go func(ipPort IpPort) {
					defer func() {
						WG.Done()
						<-threadCh
					}()
					err := Connect(ipPort.Ip, ipPort.Port)
					if err == nil {
						ResultsCh <- ipPort
					}
				}(ipPort)
			default:
				time.Sleep(time.Microsecond * 100)
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case ipPortResult, ok := <-ResultsCh:
				if !ok {
					return
				}
				SaveResult(ipPortResult.Ip, ipPortResult.Port)
				fmt.Printf("scan ip: %+v Port: %+v\n", ipPortResult.Ip, ipPortResult.Port)
			default:
				time.Sleep(time.Microsecond * 100)
			}
		}
	}(ctx)

	for _, ipObj := range ips {
		ip := ipObj.String()
		for _, port := range ports {
			ipPortsCh <- IpPort{Ip: ip, Port: port}
			WG.Add(1)
		}
	}

	WG.Wait()
	returnChan <- true
	defer func() {
		close(ipPortsCh)
		close(threadCh)
	}()

	return PrintResult()
}

func PrintResult() map[string][]int {
	result := make(map[string][]int)
	IpPortResultSyncMap.Range(func(key, value interface{}) bool {
		strKey, ok1 := key.(string)
		intSliceVal, ok2 := value.([]int)
		if ok1 && ok2 {
			result[strKey] = intSliceVal
		}
		fmt.Printf("ip:%v\n", key)
		fmt.Printf("ports: %v\n", value)
		fmt.Println(strings.Repeat("-", 20))
		return true
	})
	return result
}
