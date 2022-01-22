package player

import (
	"runtime"

	"github.com/wenchangshou2/zutil"
)

var executePath map[string]IPlayer

// GetPlayerPath 获取播放器可执行路径
func GetPlayerPath(service string) string {
	p := ""
	switch service {
	case "video":
		p, _ = zutil.GetFullPath("app/VideoPlayer/ZOOLON_VideoPlayer")
	case "web":
		p, _ = zutil.GetFullPath("app/WebPlayer/ZOOLON_WebPlayer")
	case "pdf":
		p, _ = zutil.GetFullPath("app/PDFPlayer/ZOOLON_PDFPlayer")
	case "ppt":
		p, _ = zutil.GetFullPath("app/PPTPlayer/ZOOLON_PPTPlayer")
	case "image":
		p, _ = zutil.GetFullPath("app/ImagePlayer/ZOOLON_ImagePlayer")
	default:
		return ""
	}
	if runtime.GOOS == "windows" {
		return p + ".exe"
	} else {
		return p
	}
}
func initWindow() {
	//executePath["video"] = &VideoPlayer{
	//	playPath: GetPlayerPath("video") + ".exe",
	//}
	//executePath["http"] = &HttpPlayer{
	//	playPath: GetPlayerPath("http") + ".exe",
	//}
	//executePath["pdf"] = &PdfPlayer{
	//	playPath: GetPlayerPath("pdf") + ".exe",
	//}
	//PlayerMap["ppt"] = &PptPlayer{
	//	playPath: GetPlayerPath("ppt") + ".exe",
	//}
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
