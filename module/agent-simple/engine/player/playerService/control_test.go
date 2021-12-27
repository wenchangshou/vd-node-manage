package playerService

import (
	"fmt"
	"testing"
)

func TestControlPlayer(t *testing.T) {
	s := RpcPlayerService{port: 14318}
	payload, err := s.Control("{\"Action\":\"play\"}")
	fmt.Println(payload, err)
}
func TestGetPlayerArguments(t *testing.T) {
	s := RpcPlayerService{port: 1190}
	payload, err := s.Get()
	fmt.Println(payload, err)
}
