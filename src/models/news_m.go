package models

type News struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Tags    []Tag  `json:"tags" gorm:"many2many:news_tags;"`
}