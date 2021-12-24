package playerService

import (
	"fmt"
	"testing"
)

func TestControlPlayer(t *testing.T) {
	s := RpcPlayerService{port: 8888}
	payload, err := s.Control("{\"action\":\"play\"}")
	fmt.Println(payload, err)
}
