package model

type ResourceInfo struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Uri     string `json:"uri"`
	Service string `json:"service"`
	Status  int    `json:"status"`
	Md5     string `json:"md5"`
}
type ProjectInfo struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Uri     string `json:"uri"`
	Service string `json:"service"`
	Status  int    `json:"status"`
	Startup string `json:"startup"`
	Md5     string `json:"md5"`
}

type ResourceQueryRequest struct {
	ID uint `json:"id"`
}
type ResourceQueryResponse struct {
	Code     int          `json:"code"`
	Msg      string       `json:"msg"`
	Resource ResourceInfo `json:"resource"`
}
type ProjectQueryResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Project ProjectInfo `json:"project"`
}
