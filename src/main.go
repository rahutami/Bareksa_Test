package main

import (
	"news-app/models"
	"news-app/redis_db"
	"news-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	redis_db.SetupRedis()
	routes.NewsRoutes(r)
	routes.TagsRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}