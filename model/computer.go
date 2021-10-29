package model

import (
	"time"

	"gorm.io/gorm"
)
type ComputerStatus int32
type Computer struct {
	Base
	Source         string           `gorm:"source" json:"source"`
	Switchs        string           `gorm:"switchs" json:"switchs"`
	Active         bool             `gorm:"active" json:"active"`
	Open           bool             `gorm:"open" json:"open"`
	MenuIndex      int              `gorm:"menu_index" json:"menu_index"`
	LayoutIndex    int              `gorm:"layout_index" json:"layout_index"`
	SelectedNum    string           `gorm:"selected_num" json:"selected_num" default:"[]"`
	Name           string           `gorm:"name" json:"name"`
	Ip             string           `gorm:"ip" json:"ip"`
	Mac            string           `gorm:"mac" validate:"required,mac" json:"mac"`
	HostName       string           `gorm:"hostname" json:"hostName"`
	LastOnlineTime time.Time        `gorm:"last_online_time" json:"last_online_time"`
	Screen         string           `gorm:"screen" json:"screen"`
	Resources      []Resource       `gorm:"many2many:computer_resources;" json:"_"`
	ProjectRelease []ProjectRelease ` gorm:"many2many:computer_projects;"  `
	Status int `gorm:"status" json:"status"`
}

func (Computer) TableName() string {
	return "computer"
}
func (computer Computer) Save() error {
	return DB.Model(&Computer{}).Where("id=?", computer.ID).Updates(&computer).Error
}

// AppendNewResource 添加新的资源
func (computer Computer) AppendNewResource(resource Resource) error {
	return DB.Model(resource).Association("Computers").Append(computer)
}

// AddProject 添加新的项目
func (computer Computer) AddProject(projectRelease ProjectRelease) error {
	return DB.Debug().Model(&computer).Omit("ProjectRelease.*").Association("ProjectRelease").Append(&projectRelease)
}
func (computer Computer) DeleteProject(projectRelease ProjectRelease) error {
	return DB.Debug().Model(&computer).Unscoped().Association("ProjectRelease").Delete(&projectRelease)
}
func (computer Computer) GetComputerProject(projectReleaseID string) (p *ProjectRelease, err error) {
	err = DB.Debug().Model(&computer).Where("id=?", projectReleaseID).Association("ProjectRelease").Find(&p)
	return
}
func (computer Computer)GetTasks(status TaskStatus,count int)([]Task,error){
	var res []Task
	var taskIds []string
	var taskItems []TaskItem
	err:= DB.Debug().Model(&Task{}).Joins("left join task_items on task.id=task_items.task_id").Where("computer_id=? AND task.status=?",computer.ID,status).Limit(count).Find(&res).Error
	if err!=nil{
		return nil,err
	}
	for x,task:=range res{
		res[x].TaskItems=make([]TaskItem,0)
		taskIds = append(taskIds, task.ID)
	}
	// 获取所有任务对应的子任务
	err=DB.Debug().Model(&TaskItem{}).Where("task_id=?",taskIds).Find(&taskItems).Error
	if err!=nil{
		return nil,err
	}
	// 将子任务添加到对应的任务记录当中
	for _,taskItem:=range taskItems{
		for x,task:=range res{
			if taskItem.TaskID==task.ID{
				res[x].TaskItems=append(res[x].TaskItems,taskItem )
			}
		}
	}
	return res,nil
}
func (computer Computer) ListComputerResource() (resources []Resource, err error) {
	err = DB.Debug().Model(&computer).Association("Resources").Find(&resources)
	return
}

func (computer Computer) ListComputerProject() ([]ProjectRelease, error) {
	var computerProjectList []ProjectRelease
	err := DB.Debug().Model(&computer).Association("ProjectRelease").Find(&computerProjectList)
	return computerProjectList, err
}

//DeleteResource 删除资源
func (computer Computer) DeleteResource(resource Resource) error {
	return DB.Debug().Model(&computer).Unscoped().Association("Resources").Delete(&resource)
}
func (computer *Computer) AddResource(resource Resource) error {
	return DB.Debug().Model(computer).Omit("Resources.*").Association("Resources").Append(&resource)
}
func (computer *Computer) IsExistByMac() bool {
	var client2 Computer
	err := DB.Debug().Select("mac").Where("mac = ?", computer.Mac).First(&client2).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
func (computer Computer)Heartbeat()error{
	now:=time.Now()
	return DB.Debug().Model(&computer).Update("last_online_time",now).Error
}

func (computer *Computer) UpdateByMac() error {
	data := make(map[string]interface{})
	data["ip"] = computer.Ip
	data["mac"] = computer.Mac
	data["host_name"] = computer.HostName
	data["last_online_time"] = time.Now()
	return DB.Model(&Computer{}).Where("mac=?", computer.Mac).Updates(data).Error
}
func UpdateComputerById(id string, data map[string]interface{}) error {
	return DB.Model(&Computer{}).Where("id = ?", id).Updates(data).Error
}

func (computer Computer) Create() error {
	return DB.Create(&computer).Error
}

// GetComputerByMac 通过mac地址获取用户信息
func GetComputerByMac(mac string) (Computer, error) {
	var computer Computer
	result := DB.Model(&Computer{}).Where("mac = ?", mac).First(&computer)
	return computer, result.Error
}
func GetComputerById(id interface{}) (Computer, error) {
	var computer Computer
	result := DB.Model(&Computer{}).First(&computer,"id=?",id)
	return computer, result.Error

}
func ListComputer() ([]Computer, int64) {
	var (
		computers []Computer
		total     int64
	)
	DB.Model(&Computer{}).Count(&total)
	DB.Find(&computers)
	return computers, total
}
