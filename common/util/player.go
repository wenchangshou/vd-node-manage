package util

func IsResourcePlayer(service string) bool {
	return service == "web" || service == "image" || service == "video" || service == "ppt"
}
func IsProjectPlayer(service string) bool {
	return service == "app" || service == "ue4"
}
