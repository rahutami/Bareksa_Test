package routes

import (
	controllerNews "news-app/controller/news"

	"github.com/gin-gonic/gin"
)
func NewsRoutes(r *gin.Engine){
	news := r.Group("/news")
	{
		news.GET("/", controllerNews.NewsGetAll)
		news.GET("/:id", controllerNews.NewsGetOne)
		news.POST("/", controllerNews.NewsCreate)
		news.PUT("/:id", controllerNews.NewsUpdate)
		news.DELETE("/:id", controllerNews.NewsDelete)
	}
}