package routes

import "github.com/gin-gonic/gin"

func LoadAllRoutes(r *gin.Engine) {
	CreateAuthRoute(r)
	CreateTagRoute(r)
	CreateCategoryRoute(r)
	CreateArticleRoute(r)
}
