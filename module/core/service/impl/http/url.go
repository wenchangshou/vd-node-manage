package http

import "fmt"

// GetFullUrl 获取路径
func GetFullUrl(model string) string {
	prefixUrl := "api/v1"
	url := fmt.Sprintf("http://%s:%d/%s/%s", serverIp, serverPort, prefixUrl, model)
	return url
}
