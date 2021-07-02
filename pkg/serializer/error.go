package serializer

import (
	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// AppError 应用错误，实现了error接口
type AppError struct {
	Code     int
	Msg      string
	RawError error
}

func NewError(code int, msg string, err error) AppError {
	return AppError{
		Code:     code,
		Msg:      msg,
		RawError: err,
	}
}

// Error 返回业务代码确定的可读错误信息
func (err AppError) Error() string {
	return err.Msg
}

const (
	// CodeNoPermissionErr 未授权访问
	CodeNoPermissionErr = 403
	// CodeCredentialInvalid 凭证无效
	CodeCredentialInvalid = 40001
	// CodeUploadFailed 上传出错
	CodeUploadFailed = 4002
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	// CodeCheckLogin 未登录
	CodeCheckLogin           = 401
	CodeNoFindFileErr        = 40002
	CodeNoFindProjectRelease = 40003
	CodeJsonUnMarkshalErr    = 40004
	CodeNoFoundComputerErr   = 40005
)

func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}
func Err(errCode int, msg string, err error) Response {
	// 底层错误是AppError，则尝试从AppError中获取详细信息
	if appError, ok := err.(AppError); ok {
		errCode = appError.Code
		err = appError.RawError
		msg = appError.Msg
	}
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}
