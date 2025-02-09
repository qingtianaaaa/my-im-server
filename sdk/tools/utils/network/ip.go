package network

import (
	"errors"
	"net"
)

func NewRpcRegisterIp(configIp string) (string, error) {
	registerIp := configIp
	if registerIp == "" {
		ip, err := GetLocalIp()
		if err != nil {
			return "", err
		}
		registerIp = ip
	}
	return registerIp, nil
}

func GetListenIp(configIp string) string {
	if configIp == "" {
		return "0.0.0.0"
	} else {
		return configIp
	}
}

func GetLocalIp() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var publicIp string
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil && !ipv4.IsLoopback() {
				if !ipv4.IsMulticast() {
					if !ipNet.IP.IsPrivate() && publicIp == "" {
						publicIp = ipv4.String()
					} else {
						return ipv4.String(), nil
					}
				}
			}
		}
	}
	if publicIp != "" {
		return publicIp, nil
	}
	return "", errors.New("no suitable local ip found")
}
