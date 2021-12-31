package playerService

import (
	"fmt"
	"testing"
)

func TestControlPlayer(t *testing.T) {
	s := RpcPlayerService{port: 8888}
	payload, err := s.Control("{\"Action\":\"goPage\"}")
	fmt.Println(payload, err)
}
func TestGetPlayerArguments(t *testing.T) {
	s := RpcPlayerService{port: 8888}
	payload, err := s.Get()
	fmt.Println(payload, err)
}
