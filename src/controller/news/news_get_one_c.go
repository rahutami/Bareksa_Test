package controller

import (
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

func NewsGetOne(c *gin.Context){
	var news models.News
	if err := models.DB.Where("id = ?", c.Param("id")).First(&news).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}