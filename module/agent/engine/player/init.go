package player

import (
	"runtime"

	"github.com/wenchangshou2/zutil"
)

var PlayerMap map[string]IPlayer

// 获取播放器可执行路径
func GetPlayerPath(service string) string {
	switch service {
	case "video":
		playerPath, _ := zutil.GetFullPath("app/VideoPlayer/ZOOLON_VideoPlayer")
		return playerPath
	case "http":
		playerPath, _ := zutil.GetFullPath("app/WebPlayer/ZOOLON_WebPlayer")
		return playerPath
	case "pdf":
		playerPath, _ := zutil.GetFullPath("app/PDFPlayer/ZOOLON_PDFPlayer")
		return playerPath
	case "ppt":
		playerPath, _ := zutil.GetFullPath("app/PPTPlayer/ZOOLON_PPTPlayer")
		return playerPath
	default:
		return ""
	}
}
func initWindow() {
	// PlayerMap["video"] = &VideoPlayer{
	// 	playPath: GetPlayerPath("video") + ".exe",
	// }
	// PlayerMap["http"] = &HttpPlayer{
	// 	playPath: GetPlayerPath("http") + ".exe",
	// }
	// PlayerMap["pdf"] = &PdfPlayer{
	// 	playPath: GetPlayerPath("pdf") + ".exe",
	// }
	// PlayerMap["ppt"] = &PptPlayer{
	// 	playPath: GetPlayerPath("ppt") + ".exe",
	// }
}
func initLinux() {
	// PlayerMap = make(map[string]AbstractPlayerFactory)
	// PlayerMap["video"] = &VideoPlayer{
	// 	playPath: GetPlayerPath("video"),
	// }
	// PlayerMap["http"] = &HttpPlayer{
	// 	playPath: GetPlayerPath("http"),
	// }
	// PlayerMap["pdf"] = &PdfPlayer{
	// 	playPath: GetPlayerPath("pdf"),
	// }
	// PlayerMap["ppt"] = &PptPlayer{
	// 	playPath: GetPlayerPath("ppt"),
	// }
}

func init() {
	if runtime.GOOS == "windows" {
		initWindow()
	} else {
		initLinux()
	}
}
