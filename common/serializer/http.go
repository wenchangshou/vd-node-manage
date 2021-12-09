package serializer

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

func ParamErrorMsg(filed string, tag string) string {
	// 未通过验证的表单域与中文对应
	fieldMap := map[string]string{
		"UserName": "邮箱",
		"Password": "密码",
	}
	// 未通过的规则与中文对应
	tagMap := map[string]string{
		"required": "不能为空",
		"min":      "太短",
		"max":      "太长",
		"email":    "格式不正确",
	}
	fieldVal, findField := fieldMap[filed]
	tagVal, findTag := tagMap[tag]
	if findField && findTag {
		// 返回拼接出来的错误信息
		return fieldVal + tagVal
	}
	return ""
}

func ErrorResponse(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return ParamErr(
				ParamErrorMsg(e.Field(), e.Tag()),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ParamErr("JSON类型不匹配", err)
	}
	return ParamErr("参数错误", err)
}
