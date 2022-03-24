package controller

import (
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

func TagsGetOne(c *gin.Context){
	var tag models.Tag
	if err := models.DB.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tag})
}