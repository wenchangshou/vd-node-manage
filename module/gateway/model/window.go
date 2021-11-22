package model

import "gorm.io/gorm"

type Window struct {
	Base
	Width        int    `gorm:"width"`
	Height       int    `gorm:"height"`
	X            int    `gorm:"x"`
	Y            int    `gorm:"y"`
	ExhibitionID string `gorm:"exhibition_id" json:"exhibition_id"`
	Category     string `gorm:"category" json:"category"`
	ActiveID     string `gorm:"active_id" json:"active_id"`
}

func (window *Window) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Where("exhibition_id=?", window.ID).Unscoped().Delete(&ExhibitionWindowItem{}).Error
}

func (window *Window) TableName() string {
	return "window"
}
func (window *Window) Create() (string, error) {
	if err := DB.Create(window).Error; err != nil {
		return "", err
	}
	return window.ID, nil
}
func GetExhibitionWindowByExhibitionID(id string) ([]Window, error) {
	var items []Window
	result := DB.Model(&Window{}).Where("exhibition_id=?", id).Find(&items)
	return items, result.Error
}
func DeleteExhibitionWindowByExhibitionID(id string) error {
	result := DB.Unscoped().Delete(&Window{}, "exhibition_id=?", id)
	return result.Error
}

func SetWindowActiveResourceID(wid string, aid string) error {
	return DB.Model(&Window{}).Where("id=?", wid).Update("active_id", aid).Error
}
