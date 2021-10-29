package computer

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
	"gorm.io/gorm"
)

type ListService struct {
}

func (service *ListService) List() serializer.Response {
	computer, total := model.ListComputer()
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": computer,
		},
	}
}

type UpdateService struct {
	ID       string `json:"id"`
	HostName string `json:"host_name"`
	Screen   string `json:"screen"`
	IP       string `json:"ip"`
	Mac      string `json:"mac"`
	Name     string `json:"name"`
	Status int `json:"status"`
}

func (service *UpdateService) Update() serializer.Response {
	computer, err := model.GetComputerById(service.ID)
	create := false
	if gorm.ErrRecordNotFound == err {
		computer = model.Computer{}
		computer.ID = service.ID
		create = true
	}
	if len(service.Name) > 0 {
		computer.Name = service.Name
	}
	if len(service.Screen) > 0 {
		computer.Screen = service.Screen
	}
	if len(service.HostName) > 0 {
		computer.HostName = service.HostName
	}
	if len(service.IP) > 0 {
		computer.Ip = service.IP
	}
	if len(service.Mac) > 0 {
		computer.Mac = service.Mac
	}
	if service.Status>0{
		computer.Status=service.Status
	}
	if create {
		err = computer.Create()
	} else {
		err = computer.Save()
	}
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "更新计算机信息失败", err)
	}

	return serializer.Response{}

}

type IDService struct {
	ID string `json:"id" form:"id" uri:"id"`
}
type GetDetailsForm struct {
	model.Computer
	Projects []model.ProjectRelease `json:"projects"`
}

func (service IDService) Get() serializer.Response {
	computerForm := GetDetailsForm{}
	computer, err := model.GetComputerById(service.ID)
	if err != nil || computer.ID == "" {
		return serializer.Err(serializer.CodeDBError, "获取计算机信息失败", err)
	}
	computerForm.Computer = computer
	return serializer.Response{Data: computerForm}
}
func (service IDService)Heartbeat()serializer.Response{
	computer,err:=model.GetComputerById(service.ID)
	if err!=nil{
		return serializer.Err(serializer.CodeDBError,"获取计算机实例错误",err)
	}
	err=computer.Heartbeat()
	if err!=nil{
		return serializer.Err(serializer.CodeDBError,"更新上线时间错误",err)
	}
	return serializer.Response{}
}
func (service IDService) IsRegister() serializer.Response {
	computer, err := model.GetComputerById(service.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return serializer.Err(serializer.CodeDBError, "读取计算机信息错误", err)
	}
	return serializer.Response{
		Data: computer.ID != "",
	}

}

type UpdateNameService struct {
	ID   string `json:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}

func (service *UpdateNameService) Update() serializer.Response {
	data := make(map[string]interface{})
	data["name"] = service.Name
	err := model.UpdateComputerById(service.ID, data)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "设置计算机信息失败", err)
	}
	return serializer.Response{}
}

type GetComputerService struct {
	ID string `json:"id" uri:"id" form:"id" binding:"required"`
}

func (service *GetComputerService) Get() serializer.Response {

	computer, err := model.GetComputerById(service.ID)
	if err != nil || computer.ID == "" {
		return serializer.Err(serializer.CodeDBError, "获取计算机信息失败", err)
	}
	return serializer.Response{}
}

