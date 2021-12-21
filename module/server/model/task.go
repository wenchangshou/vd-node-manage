package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/common/model"
)

type Task struct {
	Base
	Name       string     `gorm:"name" json:"name"`
	Active     bool       `gorm:"active" json:"active"`
	ComputerId string     `gorm:"computer_id" json:"computer_id"`
	Status     int        `gorm:"status" json:"status"`
	TaskItems  []TaskItem ``
}
type TaskItem struct {
	Base
	TaskID   string
	Action   uint              `gorm:"action" json:"action"`
	Status   model.EventStatus `gorm:"status" json:"status"`
	Depend   string            `gorm:"depend" json:"depend"`
	Options  string            `gorm:"options" json:"options"`
	Message  string            `gorm:"message" json:"message"`
	Schedule int               `gorm:"schedule" json:"schedule"`
}

func (taskItem *TaskItem) TableName() string {
	return "task_items"
}

func (task *Task) TableName() string {
	return "task"
}

func (task *Task) Add() (int, error) {
	if err := DB.Create(&task).Error; err != nil {
		logging.GLogger.Warn(fmt.Sprintf("添加任务项失败:%v", err))
		return -1, err
	}
	return task.ID, nil
}

// AddTaskItem 添加任务项
func AddTaskItem(taskId string, action uint, options map[string]interface{}, _ bool, depend string) (TaskItem, error) {
	jsonStr, _ := json.Marshal(options)
	taskItem := TaskItem{
		TaskID:   taskId,
		Action:   action,
		Options:  string(jsonStr),
		Schedule: 0,
		Status:   model.Initializes,
		Depend:   depend,
	}
	result := DB.Model(&TaskItem{}).Create(&taskItem)
	return taskItem, result.Error
}

// AddTask 添加任务
func AddTask(name string, computerID string) (Task, error) {
	task := Task{
		Name:       name,
		Status:     0,
		ComputerId: computerID,
	}
	result := DB.Model(&Task{}).Create(&task)
	return task, result.Error
}

func GetTaskListByCid(computerId string) ([]Task, error) {
	var tasks []Task
	result := DB.Debug().Model(&Task{}).Where("computer_id = ?", computerId).Find(&tasks)
	return tasks, result.Error
}
func GetTaskListByComputerID(Page int, size int, orderBy string, conditions map[string]string, _ map[string]string) ([]Task, int64) {
	var res []Task
	var total int64
	tx := DB.Model(&Task{})
	if orderBy != "" {
		tx = tx.Order(orderBy)
	}
	for k, v := range conditions {
		tx = tx.Where(k+" = ?", v)
	}
	tx.Count(&total)
	tx.Debug().Limit(size).Offset((Page - 1) * size).Association("TaskItems").Find(&res)
	return res, total
}
func GetTaskListByCidFilterStatus(computerId string, status int) ([]Task, error) {
	var tasks []Task
	result := DB.Debug().Model(&Task{}).Where("computer_id = ? AND status = ?", computerId, status).Find(&tasks)
	return tasks, result.Error
}

func SetTaskItemStatus(taskId string, status uint, msg string) error {
	result := DB.Model(&TaskItem{}).Where("id like ?", taskId).Updates(map[string]interface{}{"status": status, "message": msg})
	return result.Error

}

// GetTasks 获取所有任务项
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
			search += k + " like '%" + v + "%' OR "
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	err := tx.Debug().Limit(size).Offset((page - 1) * size).Association("TaskItems").Find(&res)
	if err != nil {
		return nil, 0
	}
	return res, total
}

// GetTaskItemById 通过id获取任务项
func GetTaskItemById(id string) ([]TaskItem, int64) {
	var res []TaskItem
	var total int64
	result := DB.Model(&TaskItem{}).Where("task_id = ?", id).Find(&res)
	result.Count(&total)
	return res, total
}

// GetTasksByDependID 通过依赖获取任务项
func GetTasksByDependID(id string) ([]Task, int64) {
	var res []Task
	var total int64
	result := DB.Model(&Task{}).Where("depend = ?", id).Find(&res)
	result.Count(&total)
	return res, total
}

func ListTask(Page int, size int, orderBy string, conditions map[string]string, _ map[string]string) ([]Task, int64) {
	var res []Task
	var total int64
	tx := DB.Model(&Task{})
	if orderBy != "" {
		tx = tx.Order(orderBy)
	}
	for k, v := range conditions {
		tx = tx.Where(k+" = ?", v)
	}
	tx.Count(&total)
	tx.Debug().Limit(size).Offset((Page - 1) * size).Preload("TaskItems").Find(&res)
	return res, total
}
func UpdateTaskStatusByIds(ids []string, status int) error {
	return DB.Model(&Task{}).Where("ID=?", ids).Update("status", status).Error
}
func UpdateTaskItemStatusByIds(ids []string, status int) error {
	return DB.Model(&TaskItem{}).Where("ID=?", ids).Update("status", status).Error
}
