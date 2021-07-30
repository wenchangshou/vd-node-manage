package model

import (
	"encoding/json"
	"fmt"
	"strings"

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
	Name       string `gorm:"name"`
	Active     bool   `gorm:"active"`
	ComputerId int    `gorm:"computer_id"`
	Status     int    `gorm:"status"`
}
type TaskItem struct {
	gorm.Model
	TaskID   uint   `gorm:"task_id"`
	Action   uint   `gorm:"action"`
	Status   int    `gorm:"status"`
	Depend   int    `gorm:"depend"`
	Options  string `gorm:"options"`
	Message  string `gorm:"message"`
	Schedule int    `gorm:"schedule"`
}

func (taskItem *TaskItem) TableName() string {
	return "task_items"
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

// AddTaskItem 添加任务项
func AddTaskItem(taskId uint, action uint, options map[string]interface{}, active bool, depend uint) (TaskItem, error) {
	jsonStr, _ := json.Marshal(options)
	taskItem := TaskItem{
		TaskID:   taskId,
		Action:   action,
		Options:  string(jsonStr),
		Schedule: 0,
		Status:   Initializes,
		Depend:   int(depend),
	}
	result := DB.Model(&TaskItem{}).Create(&taskItem)
	return taskItem, result.Error
}

func AddTask(name string, computerID uint) (Task, error) {
	task := Task{
		Name:       name,
		Status:     0,
		ComputerId: int(computerID),
	}
	result := DB.Model(&Task{}).Create(&task)
	return task, result.Error
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
func SetTaskItemStatus(taskId uint, status uint, msg string) error {
	result := DB.Model(&TaskItem{}).Where("id = ?", taskId).Updates(map[string]interface{}{"status": status, "message": msg})
	return result.Error

}

func GetTasks(page int, size int, orderBy string, conditions map[string]string, searches map[string]string) ([]Task, int64) {
	var res []Task
	var total int64
	tx := DB.Model(&Task{})
	if orderBy != "" {
		tx = tx.Order(orderBy)
	}
	for k, v := range conditions {
		tx = tx.Where(k+" = ?", v)
	}
	if len(searches) > 0 {
		search := ""
		for k, v := range searches {
			search += (k + " like '%" + v + "%' OR ")
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	tx.Debug().Limit(size).Offset((page - 1) * size).Find(&res)
	return res, total
}
func GetTaskItemById(id int) ([]TaskItem, int64) {
	var res []TaskItem
	var total int64
	result := DB.Model(&TaskItem{}).Where("task_id = ?", id).Find(&res)
	result.Count(&total)
	return res, total
}
func GetTasksByDependID(id int) ([]Task, int64) {

	var res []Task
	var total int64
	result := DB.Model(&Task{}).Where("depend = ?", id).Find(&res)
	result.Count(&total)
	return res, total
}
