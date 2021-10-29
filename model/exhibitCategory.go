package model

type ExhibitionCategory struct {
	Base
	Name       string
	ComputerID string `gorm:"computer_id" json:"computer_id"`
	Level      int    `gorm:"level" json:"level"`
}

func (exhibitionCategory *ExhibitionCategory) TableName() string {
	return "exhition_category"
}
func (exhibitionCategory *ExhibitionCategory) Create() (string, error) {
	if err := DB.Create(exhibitionCategory).Error; err != nil {
		return "", err
	}
	return exhibitionCategory.ID, nil
}

//GetComputerExhibtionCatetory 获取指定计算机的展项类别
func GetComputerExhibtionCatetory(cid string) ([]ExhibitionCategory, error) {
	var categorys []ExhibitionCategory
	result := DB.Where("computer_id=?", cid).Find(&categorys)
	return categorys, result.Error
}
