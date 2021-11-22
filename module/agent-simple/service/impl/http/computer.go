package http

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/ao"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	"net/http"
	"os"
)

type ComputerHttpService struct {
	ID   string
	Ip   string
	Port uint
	GetUrlStruct
	Address string
}

// ReportServiceInfo 上报信息
func (service ComputerHttpService) ReportServiceInfo(id string, ip string, mac string, name string) error {
	rtu := e.HttpBaseData{}
	client := resty.New()
	rUrl := service.GetUrl(service.Address, "computer")
	resp, err := client.R().SetBody(map[string]interface{}{
		"host_name": name,
		"mac":       mac,
		"ip":        ip,
		"id":        id,
	}).SetResult(&rtu).Put(rUrl)
	if err != nil {
		return errors.Wrap(err, "调用更新计算机信息失败")
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(fmt.Sprintf("调用更新计算机信息失败,返回代码:%d", resp.StatusCode()))
	}
	if rtu.Code != 0 {
		return errors.New("调用更新计算机信息失败，错误信息:" + rtu.Msg)
	}
	return nil
}

func NewComputerHttpService(id string, address string) *ComputerHttpService {
	return &ComputerHttpService{
		ID:      id,
		Address: address,
	}
}

// IsRegister 是否注册
func (service ComputerHttpService) IsRegister() (bool, error) {
	rtu := ao.ComputerRegisterResultForm{}
	client := resty.Client{}
	requestUrl := GetFullUrl(fmt.Sprintf("computer/%s/register", service.ID))
	response, err := client.R().SetResult(&rtu).Get(requestUrl)
	if err != nil {
		return false, err
	}
	if response.StatusCode() != http.StatusOK {
		return false, errors.New("获取注册状态失败")
	}
	if rtu.Code != 0 {
		return false, errors.New(rtu.Msg)
	}
	return rtu.Data, nil
}
func (service ComputerHttpService) Report() error {
	client := resty.New()
	rtu := e.HttpBaseData{}
	requestUri := service.GetUrl(service.Address, "computer/"+service.ID+"/report")
	name, err := os.Hostname()
	if err != nil {
		return errors.Wrap(err, "获取计算机主机名失败")
	}
	resp, err := client.R().SetBody(map[string]interface{}{
		"id":   service.ID,
		"name": name,
	}).Post(requestUri)
	if err != nil {
		return errors.Wrap(err, "更新计算机信息失败")
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New("请示更新计算机信息错误")
	}
	if rtu.Code != 0 {
		return errors.New("更新计算机信息失败，错误信息:" + rtu.Msg)
	}
	return nil

}

func (service ComputerHttpService) Heartbeat() error {
	client := resty.New()
	rtu := e.HttpBaseData{}
	requestUrl := service.GetUrl(service.Address, "computer/"+service.ID+"/heartbeat")
	resp, err := client.R().SetResult(&rtu).Get(requestUrl)
	if err != nil {
		return errors.Wrap(err, "更新心跳数据失败")
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(fmt.Sprintf("请示心跳失败，返回的错误代码(%d)", resp.StatusCode()))
	}
	if rtu.Code != 0 {
		return errors.New("请示心跳失败,返回的错误消息:" + rtu.Msg)
	}
	return nil

}

//AddComputerProject 添加计算机项目
func (service ComputerHttpService) AddComputerProject(projectReleaseId string) error {
	rtu := e.HttpBaseData{}
	requestUri := fmt.Sprintf("computer/%s/projectRelease/%s", service.ID, projectReleaseId)
	requestUri = service.GetUrl(service.Address, requestUri)
	client := resty.New()
	req, err := client.R().SetResult(&rtu).Post(requestUri)
	if req.StatusCode() != http.StatusOK || req.Error() != nil || err != nil {
		return errors.New("添加计算机项目失败")
	}
	if rtu.Code != 0 {
		return errors.New("添加计算机项目失败,错误信息:" + rtu.Msg)
	}
	return nil
}

// DeleteComputerProject 删除计算机项目
func (service ComputerHttpService) DeleteComputerProject(projectReleaseId string) error {
	rtu := e.HttpBaseData{}
	requestUri := fmt.Sprintf("computer/%s/projectRelease/%s", service.ID, projectReleaseId)
	requestUri = GetFullUrl(requestUri)
	client := resty.Client{}
	req, err := client.R().SetResult(&rtu).Delete(requestUri)
	if req.StatusCode() != http.StatusOK || req.Error() != nil || err != nil {
		return errors.New("删除计算机项目失败")
	}
	if rtu.Code != 0 {
		return errors.New("删除计算机项目失败,错误信息:" + rtu.Msg)
	}
	return nil
}

// DeleteComputerResource 删除计算机资源
func (service ComputerHttpService) DeleteComputerResource(sid string) error {
	rtu := e.HttpBaseData{}
	requestUri := fmt.Sprintf("computer/%s/resource/%s", service.ID, sid)
	requestUri = GetFullUrl(requestUri)
	client := resty.New()
	req, err := client.R().SetResult(&rtu).Delete(requestUri)
	if req.StatusCode() != http.StatusOK || req.Error() != nil || err != nil {
		return errors.New("删除计算机资源失败")
	}
	if rtu.Code != 0 {
		return errors.New("删除计算机资源失败,错误信息:" + rtu.Msg)
	}
	return nil
}

// AddComputerResource 添加计算机资源
func (service ComputerHttpService) AddComputerResource(sid string) error {
	rtu := e.HttpBaseData{}
	requestUri := fmt.Sprintf("computer/%s/resource/%s", service.ID, sid)
	requestUri = service.GetUrl(service.Address, requestUri)
	client := resty.New()
	req, err := client.R().SetResult(&rtu).Post(requestUri)
	if req.StatusCode() != http.StatusOK || req.Error() != nil || err != nil {
		return errors.New("添加计算机资源失败")
	}
	if rtu.Code != 0 {
		return errors.New("添加计算机资源失败,错误信息:" + rtu.Msg)
	}
	return nil
}
