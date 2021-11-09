package e

//HttpBaseData http返回标准的结构
type HttpBaseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
