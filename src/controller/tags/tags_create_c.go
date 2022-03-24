package controller

import (
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

type TagCreateInput struct {
	TagName string `json:"tagname" binding:"required"`
}

func TagsCreate(c *gin.Context){
	var input TagCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{TagName: input.TagName, News: []models.News{}}
	models.DB.Create(&tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}