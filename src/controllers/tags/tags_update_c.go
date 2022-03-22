package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsUpdate(c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"news": "update",
	})
}