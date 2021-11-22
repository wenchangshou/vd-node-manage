package zebus

// DirectoryInfoRequestForm 目录信息请求表单
type DirectoryInfoRequestForm struct {
	Action string `json:"action"`
	Dir    string `json:"dir"`
}

// DirectoryInfo 目录信息
type DirectoryInfo struct {
	Name string
	Type string `json:"type"`
}

type DirectoryInfoResponseForm struct {
	Dir  string          `json:"dir"`
	Node []DirectoryInfo `json:"node"`
}
