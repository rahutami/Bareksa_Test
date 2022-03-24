package controller

import (
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

type UpdateNewsInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Tags []uint `json:"tags" binding:"required"`
}

func NewsUpdate(c *gin.Context){
	var news models.News
	if err := models.DB.Where("id = ?", c.Param("id")).First(&news).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateNewsInput
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

	models.DB.Model(&news).Updates(input)
	models.DB.Model(&news).Update("Tags", tags)

	c.JSON(http.StatusOK, gin.H{"data": news})
}