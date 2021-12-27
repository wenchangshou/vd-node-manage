package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg2/e"
	"net/http"
)

// GetExternIp  获取当前的计算机ip
func GetExternIp() (ip string, err error) {
	rtu := e.HttpBaseData{}
	fullUrl := GetFullUrl("system/extranet")
	client := resty.New()
	resp, err := client.R().SetResult(&rtu).Get(fullUrl)
	if resp.StatusCode() != http.StatusOK {
		return "", errors.New("请示获取计算机ip接口失败")
	}
	if err != nil {
		return "", errors.Wrap(err, "获取计算机外部ip失败")
	}
	if rtu.Code != 0 {
		return "", errors.New("获取计算机外部ip失败，错误消息:" + rtu.Msg)
	}
	return rtu.Data.(string), nil
}
