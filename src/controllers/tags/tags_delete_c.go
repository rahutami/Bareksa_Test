package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsDelete(c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"news": "delete",
	})
}