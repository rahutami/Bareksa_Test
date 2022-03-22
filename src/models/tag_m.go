package models

type Tag struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	TagName string `json:"tagname"`
}