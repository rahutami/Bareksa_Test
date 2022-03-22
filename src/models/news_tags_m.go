package models

type NewsTags struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	IDNews string `json:"id_news"`
	IDTag  string `json:"id_tag"`
}