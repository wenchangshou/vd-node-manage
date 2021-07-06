package zutil

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// DotPathToStandardPath 将","分割的路径转换为标准路径
func DotPathToStandardPath(path string) string {
	return "/" + strings.Replace(path, ",", "/", -1)
}

// FillSlash 给路径补全`/`
func FillSlash(path string) string {
	if path == "/" {
		return path
	}
	return path + "/"
}

// RemoveSlash 移除路径最后的`/`
func RemoveSlash(path string) string {
	if len(path) > 1 {
		return strings.TrimSuffix(path, "/")
	}
	return path
}

// SplitPath 分割路径为列表
func SplitPath(path string) []string {
	if len(path) == 0 || path[0] != '/' {
		return []string{}
	}

	if path == "/" {
		return []string{"/"}
	}

	pathSplit := strings.Split(path, "/")
	pathSplit[0] = "/"
	return pathSplit
}

// FormSlash 将path中的反斜杠'\'替换为'/'
func FormSlash(old string) string {
	return path.Clean(strings.ReplaceAll(old, "\\", "/"))
}

// RelativePath 获取相对可执行文件的路径
func RelativePath(name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), name)
}

/* ==========
	 验证器
   ==========
*/

// 文件/路径名保留字符
var reservedCharacter = []string{"\\", "?", "*", "<", "\"", ":", ">", "/", "|"}

// IsInExtensionList 返回文件的扩展名是否在给定的列表范围内
func IsInExtensionList(extList []string, fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	// 无扩展名时
	if len(ext) == 0 {
		return false
	}

	if ContainsString(extList, ext[1:]) {
		return true
	}

	return false
}
