package models

type Tag struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	TagName string `json:"tagname"`
	News    []News `json:"news" gorm:"many2many:news_tags;"`
}