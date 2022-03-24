package controller

import (
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

type TagUpdateInput struct {
	TagName string `json:"tagname"`
}

func TagsUpdate(c *gin.Context){
	var tags models.Tag
	if err := models.DB.Where("id = ?", c.Param("id")).First(&tags).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input TagUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&tags).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": tags})
}