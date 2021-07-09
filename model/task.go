package model

import (
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"gorm.io/gorm"
)

const (
	Initializes = iota
	Progress
	Done
	Error
)
const (
	InstallProjectAction = iota
	InstallResourceAction
	UpgradeProjectAction
	DeleteResource
	DeleteProject
)

type Task struct {
	gorm.Model
	ComputerId int    `gorm:"computer_id"`
	Options    string `gorm:"options"`
	Action     uint   `gorm:"action"`
	Status     int    `gorm:"status"`
	Depend     int    `gorm:"depend"`
	Schedule   int    `gorm:"schedule"`
	Active     bool   `gorm:"active"`
}

func (task *Task) TableName() string {
	return "task"
}

func (task *Task) Add() (uint, error) {
	if err := DB.Create(&task).Error; err != nil {
		logging.G_Logger.Warn(fmt.Sprintf("添加任务项失败:%v", err))
		return 0, err
	}
	return task.ID, nil
}
func AddTasks(computers []uint, action uint, options string, active bool) (interface{}, error) {
	tx := DB.Begin()
	ids := make([]uint, 0)
	for computer := range computers {
		task := Task{
			Options:    options,
			Action:     action,
			Status:     Initializes,
			ComputerId: computer,
			Active:     active,
		}

		id, err := task.Add()
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
func NewTask(computerid int, action int, option string, depend int) Task {
	return Task{
		ComputerId: computerid,
		Action:     uint(action),
		Status:     Initializes,
		Options:    option,
		Depend:     depend,
	}
}

//GetPendingTaskByComputerId 获取待处理的任务
func GetPendingTaskByComputerId(computerId int) ([]Task, error) {
	var tasks []Task
	result := DB.Model(&Task{}).Where("computer_id = ? AND status != ?", computerId, Done).Find(&tasks)
	return tasks, result.Error
}
func GetTaskListByCid(computerId int) ([]Task, error) {
	var tasks []Task
	result := DB.Debug().Model(&Task{}).Where("computer_id = ?", computerId).Find(&tasks)
	return tasks, result.Error
}
func GetTaskListByCidFilterStatus(computerId int, status int) ([]Task, error) {
	var tasks []Task
	result := DB.Debug().Model(&Task{}).Where("computer_id = ? AND status = ?", computerId, status).Find(&tasks)
	return tasks, result.Error
}

// SetTaskStatus 设置任务的状态
func SetTaskStatus(taskId uint, status uint) error {
	result := DB.Model(&Task{}).Where("id = ? ", taskId).Update("status", status)
	return result.Error
}
