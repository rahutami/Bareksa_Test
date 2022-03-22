package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsCreate(c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"news": "create",
	})
}