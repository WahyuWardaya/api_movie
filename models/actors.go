package models

type Actors struct {
	ID     uint     `gorm:"column:id_actors;primary_key" json:"id"`
	Name   string   `gorm:"column:name_actors" json:"name_actors"`
	
}

func (Actors) TableName() string {
	return "actors"
}