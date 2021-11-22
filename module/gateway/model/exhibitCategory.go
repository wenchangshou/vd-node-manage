package model

type ExhibitionCategory struct {
	Base
	Name       string
	ComputerID string `gorm:"computer_id" json:"computer_id"`
	Level      int    `gorm:"level" json:"level"`
}

func (exhibitionCategory *ExhibitionCategory) TableName() string {
	return "exhibition_category"
}
func (exhibitionCategory *ExhibitionCategory) Create() (string, error) {
	if err := DB.Create(exhibitionCategory).Error; err != nil {
		return "", err
	}
	return exhibitionCategory.ID, nil
}

// GetComputerExhibitionCategory  获取指定计算机的展项类别
func GetComputerExhibitionCategory(cid string) ([]ExhibitionCategory, error) {
	var categories []ExhibitionCategory
	result := DB.Where("computer_id=?", cid).Find(&categories)
	return categories, result.Error
}
