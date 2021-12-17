package model

type ExhibitionWindowItem struct {
	Base
	ExhibitionID string   `gorm:"exhibition_id" json:"exhibition_id"`
	WindowID     string   `gorm:"window_id" json:"window_id"`
	ResourceID   string   `gorm:"resource_id" json:"resource_id"`
	ProjectID    string   `gorm:"project_id" json:"project_id"`
	ModuleID     string   `gorm:"module_id" json:"module_id"`
	Window       Window   `json:"window" gorm:"foreignkey:WindowID"`
	Resource     Resource `json:"resource" gorm:"foreignkey:ResourceID"`
	Project      Project  `json:"project" gorm:"foreignkey:ProjectID"`
	Module       Module   `json:"module" gorm:"foreignkey:ModuleID"`
}

func (exhibitionWindowItem *ExhibitionWindowItem) TableName() string {
	return "exhibition_window_item"
}

func (exhibitionWindowItem *ExhibitionWindowItem) Create() error {
	return DB.Model(&ExhibitionWindowItem{}).Create(&exhibitionWindowItem).Error
}

func GetExhibitionDetailsByExhibitionID(id string) ([]ExhibitionWindowItem, error) {
	var items []ExhibitionWindowItem
	err := DB.Debug().Model(&ExhibitionWindowItem{}).
		Preload("Project").Preload("Resource").
		Preload("Window").Where("exhibition_id=?", id).
		Find(&items).Error
	return items, err
}

func DeleteExhibitionWindowItemByExhibitionID(id string) error {
	return DB.Debug().Unscoped().Delete(&ExhibitionWindowItem{}, "exhibition_id=?", id).Error
}

func GetSpecifiedModuleExhibition(id string) ([]string, error) {
	ids := make([]string, 0)
	var items []ExhibitionWindowItem
	result := DB.Model(&ExhibitionWindowItem{}).Where("module_id=?", id).Find(&items)
	if result.Error != nil {
		return ids, result.Error
	}
	for i := range items {
		ids = append(ids, items[i].ExhibitionID)
	}
	return ids, nil
}
