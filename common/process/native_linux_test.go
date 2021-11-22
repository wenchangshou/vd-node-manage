package process

import (
	"fmt"
	"testing"
)

func TestGetProcessIdByName(t *testing.T) {
	pid, err := GetProcessIdByName("zebus")
	fmt.Println(pid, err)
}
