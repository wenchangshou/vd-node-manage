package model

type ComputerExhibition struct {
	Base
	ComputerID   string `gorm:"computer_id" json:"computer_id"`
	CategoryID   string `gorm:"category_id" json:"category_id"`
	ExhibitionID string `gorm:"exhibition_id" json:"exhibition_id"`
}

func (computerExhibition *ComputerExhibition) TableName() string {
	return "computer_exhibition"
}
