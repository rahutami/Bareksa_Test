package controller

import (
	"context"
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

type CreateNewsInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status"`
	Tags []uint `json:"tags"`
}

type NewsResponse struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Tags []models.Tag `json:"tags"`
}
func NewsCreate(c *gin.Context){
	var input CreateNewsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tags := []models.Tag{}
	if err := models.DB.Find(&tags, input.Tags).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag record not found!"})
			return
	}

	if(len(tags) != len(input.Tags)){	
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag record not found!"})
		return
	}
	news := models.News{Title: input.Title, Content: input.Content, Status: "Draft", Tags: tags}
	models.DB.Create(&news)

	for _, tag := range tags {
		newsAssociation := 
		models.DB.Model(&tag).Association("News")
		if (newsAssociation == nil){
			newsAssociation.Replace([]models.News{news})
		} else {
			newsAssociation.Append(&news)
		}
		models.DB.Model(&tag).Update("News", &newsAssociation)
	}
	c.JSON(http.StatusOK, gin.H{"data": news})
}