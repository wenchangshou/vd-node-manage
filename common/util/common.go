package util

import "strings"

// Replace 根据替换表执行批量替换
func Replace(table map[string]string, s string) string {
	for key, value := range table {
		s = strings.ReplaceAll(s, key, value)
	}
	return s
}
