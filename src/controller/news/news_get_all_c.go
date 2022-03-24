package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news-app/models"
	"news-app/redis_db"

	"github.com/gin-gonic/gin"
)

func NewsGetAll(c *gin.Context){
	status := c.Request.URL.Query().Get("status")
	val, err := redis_db.RDB.Get(ctx, "newsAll").Result()
	if err == nil && status == "" {
		var result []models.News
		json.Unmarshal([]byte(val), &result)
		c.JSON(http.StatusOK, gin.H{"data": result, "source": "redis"})
	} else {
		var news []models.News
		if (status != ""){
			models.DB.Find(&news, "status = ?", status)
		}else{
			models.DB.Find(&news)
			json, err := json.Marshal(news)

			if err != nil {
				fmt.Println(err)
			}

			err = redis_db.RDB.Set(ctx, "newsAll", json, 0).Err()
			if err != nil {
				fmt.Println(err)
			}
		}

		c.JSON(http.StatusOK, gin.H{"data": news, "source": "database"})
	}
}