package http

var (
	serverIp   string
	serverPort int
)

func InitService(ip string, port int) {
	serverIp = ip
	serverPort = port
}
