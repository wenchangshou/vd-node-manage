package zebus

import (
	"fmt"
	"testing"
)

func TestGetClients(t *testing.T) {
	InitZebus("192.168.0.223", 9191, 8181)
	form, err := G_Zebus.GetClients()
	fmt.Println(form, err)
}
func TestPutMessage(t *testing.T) {
	InitZebus("192.168.0.222", 9191, 8181)
}