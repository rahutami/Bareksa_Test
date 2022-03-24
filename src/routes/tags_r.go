package routes

import (
	controllerTags "news-app/controller/tags"

	"github.com/gin-gonic/gin"
)
func TagsRoutes(r *gin.Engine){
	tags := r.Group("/tags")
	{
		tags.GET("/", controllerTags.TagsGetAll)
		tags.GET("/:id", controllerTags.TagsGetOne)
		tags.POST("/", controllerTags.TagsCreate)
		tags.PUT("/:id", controllerTags.TagsUpdate)
		tags.DELETE("/:id", controllerTags.TagsDelete)
	}
}