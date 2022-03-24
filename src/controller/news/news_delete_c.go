package controller

import (
	"net/http"
	"news-app/models"

	"github.com/gin-gonic/gin"
)

type DeletedNews struct {
	Status  string `json:"status"`
}

func NewsDelete(c *gin.Context){
	var news models.News
	var deleteUpdate DeletedNews
	if err := models.DB.Where("id = ?", c.Param("id")).First(&news).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	deleteUpdate.Status = "Deleted"
	models.DB.Model(&news).Updates(deleteUpdate)

	c.JSON(http.StatusOK, gin.H{"data": true})
}