package http

import (
	"fmt"
	"testing"
)

func TestGetExternIp(t *testing.T) {
	InitService("192.168.10.121", 8000)
	ip, err := GetExternIp()
	fmt.Println("ip", ip, err)
}
