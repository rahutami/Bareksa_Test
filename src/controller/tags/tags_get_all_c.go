package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"news-app/models"
	"news-app/redis_db"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func TagsGetAll(c *gin.Context){
	val, err := redis_db.RDB.Get(ctx, "tagsAll").Result()
	if err == nil{
		var result []models.News
		json.Unmarshal([]byte(val), &result)
		c.JSON(http.StatusOK, gin.H{"data": result, "source": "redis"})
	} else {
		var tags []models.Tag
		models.DB.Find(&tags)

    	json, err := json.Marshal(tags)

		if err != nil {
			fmt.Println(err)
		}

		err = redis_db.RDB.Set(ctx, "tagsAll", json, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": tags, "source": "database"})
	}
}