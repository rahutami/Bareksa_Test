package controller

import (
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

func TagsDelete(c *gin.Context){
	var tags models.Tag
	if err := models.DB.Where("id = ?", c.Param("id")).First(&tags).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&tags)

	c.JSON(http.StatusOK, gin.H{"data": true})
}