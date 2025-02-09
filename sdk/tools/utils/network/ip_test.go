package network

import (
	"fmt"
	"net"
	"testing"
)

func Test(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err : ", err)
	}
	for _, ifaces := range interfaces {
		fmt.Println("--------")
		// fmt.Println(ifaces)
		// fmt.Println(ifaces.Addrs())
		addrs, _ := ifaces.Addrs()
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}
			ip4 := ipNet.IP.To4()
			fmt.Println(ip4)
		}
	}
}

func Test2(t *testing.T) {
	res, err := NewRpcRegisterIp("")
	if err != nil {
		fmt.Println(" err : ", err)
	}
	fmt.Println(res)
}
