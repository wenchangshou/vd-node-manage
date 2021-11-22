package http

import "fmt"

type GetUrlStruct struct {
	Ip   string
	Port int
}

func (s GetUrlStruct) GetUrl(address string, model string) string {
	prefixUrl := "api/v1"
	url := fmt.Sprintf("http://%s/%s/%s", address, prefixUrl, model)
	return url

}
