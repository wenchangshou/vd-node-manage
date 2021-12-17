package util

import (
	"errors"
	"net"
	"strings"
)

func GetMacByIp(ip string) (mac string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, netInterface := range netInterfaces {
		addrs, _ := netInterface.Addrs()
		macAddr := netInterface.HardwareAddr.String()
		if macAddr == "" {
			continue
		}
		for _, addr := range addrs {
			if strings.HasPrefix(addr.String(), ip) {
				return macAddr, nil
			}
		}
	}
	return "", errors.New("没有找到对应的mac记录")
}
