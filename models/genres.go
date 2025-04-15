package models

type Genres struct {
	ID     uint     `gorm:"column:id_genres;primary_key" json:"id"`
	Genres string   `gorm:"column:genres" json:"genres"`
}

func (Genres) TableName() string {
	return "genres"
}