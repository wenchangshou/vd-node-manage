package playerService

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type HttpPlayerService struct {
	port int `json:"port"`
}

func (svr HttpPlayerService) Ping() (bool, error) {
	client := resty.New().SetTimeout(500 * time.Millisecond)
	resp, err := client.R().Get(fmt.Sprintf("http://localhost:%d/ping", svr.port))
	if err != nil {
		return false, err
	}
	return string(resp.Body()) == "pong", nil

}
func (svr HttpPlayerService) Control(body string) (reply string, err error) {
	var (
		resp *resty.Response
	)
	client := resty.New().SetTimeout(500 * time.Millisecond)
	if resp, err = client.R().SetBody(body).Post(fmt.Sprintf("http://localhost:%d/control", svr.port)); err != nil {
		return "", err
	}
	return string(resp.Body()), nil
}
