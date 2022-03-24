package models

type NewsTags struct {
	ID     uint `json:"id" gorm:"primary_key"`
	IDNews uint `json:"id_news"`
	IDTag  uint `json:"id_tag"`
}