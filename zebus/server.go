package zebus

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

var (
	G_Zebus *Zebus
)

type Zebus struct {
	Ip       string
	HttpPort int
	WsPort   int
}
type OnlineComputerInfo struct {
	Ip       string
	Server   []string
	Config   map[string]interface{}
	Resource []string
}
type GetClientForm struct {
	Offline []string             `json:"offline" `
	Server  []string             `json:"server"`
	Online  []OnlineComputerInfo `json:"online"`
}

func (service GetClientForm) IsExistServer(ip string, server string) bool {
	for _, computerInfo := range service.Online {
		if computerInfo.Ip == ip {
			for _, computerServer := range computerInfo.Server {
				if strings.EqualFold(computerServer, server) {
					return true
				}
			}
		}
	}
	return false
}

func (zebus *Zebus) GetClients() (*GetClientForm, error) {
	var form *GetClientForm
	client := resty.New()
	resp, err := client.R().SetResult(&form).
		Post(fmt.Sprintf("http://%s:%d/getClients", zebus.Ip, zebus.HttpPort))
	if err != nil {
		return nil, err
	}
	fmt.Println("resp", resp)
	fmt.Println(form)
	return form, nil
}
func (zebus *Zebus) PutV2(topic string, body string, _ int, _ int, result interface{}) error {
	client := resty.New()
	dstAddr := fmt.Sprintf("http://%s:%d/pubV2?topic=%s", zebus.Ip, zebus.HttpPort, topic)
	resp, err := client.R().SetResult(&result).
		SetBody(body).
		SetResult(&result).
		Post(dstAddr)
	fmt.Println("resp", resp)
	return err
}
func InitZebus(Ip string, HttpPort, WsPort int) {
	G_Zebus = &Zebus{
		Ip:       Ip,
		HttpPort: HttpPort,
		WsPort:   WsPort,
	}
}
