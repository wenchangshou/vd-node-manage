package HttpService

import "fmt"

type GetUrlStruct struct{
	Ip string
	Port int
}
func (s GetUrlStruct)GetUrl(ip string,port int,model string)string{
	prefixUrl := "api/v1"
	url := fmt.Sprintf("http://%s:%d/%s/%s", ip,port,prefixUrl, model)
	return url

}
