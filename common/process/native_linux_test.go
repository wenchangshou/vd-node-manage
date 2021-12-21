package process

import (
	"fmt"
	"testing"
)

func TestGetProcessIdByName(_ *testing.T) {
	pid, err := GetProcessIdByName("zebus")
	fmt.Println(pid, err)
}
