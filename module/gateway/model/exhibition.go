package model

type Exhibition struct {
	Base
	CategoryID string `gorm:"category_id"`
	Name       string `gorm:"name" json:"name"`
	Level      int    `gorm:"level" json:"level"`
	Encryption bool   `gorm:"encryption" json:"encryption"`
	Password   string `gorm:"password" json:"password"`
	Control    string `gorm:"control" json:"control"`
}

func (exhibition *Exhibition) TableName() string {
	return "exhibition"
}

func (exhibition *Exhibition) Create() (string, error) {
	if err := DB.Create(exhibition).Error; err != nil {
		return "", err
	}
	return exhibition.ID, nil
}
func (exhibition *Exhibition) Update(val map[string]interface{}) error {
	return DB.Model(exhibition).Updates(val).Error
}

//GetExhibitionByCategory 通过类别获取展项
func GetExhibitionByCategory(categoryId string) ([]Exhibition, error) {
	var exhibitionList []Exhibition
	result := DB.Model(&Exhibition{}).Where("category_id=?", categoryId).Find(&exhibitionList)
	return exhibitionList, result.Error
}

func GetExhibitionByID(id string) (exhibition Exhibition, err error) {
	result := DB.Model(&Exhibition{}).Where("id=?", id).First(&exhibition)
	return exhibition, result.Error
}

func DeleteExhibitionByID(id string) error {
	result := DB.Unscoped().Delete(&Exhibition{}, "id=?", id)
	return result.Error
}
