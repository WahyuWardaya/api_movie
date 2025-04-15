package models

type Directors struct {
	ID     uint     `gorm:"column:id_directors;primary_key" json:"id"`
	Name   string   `gorm:"column:name_director" json:"name_director"`
	
}

func (Directors) TableName() string {
	return "directors"
}