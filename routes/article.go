package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticleRoute(r *gin.Engine) {
	articleRouter := r.Group("api/article")
	{
		// 获取文章
		articleRouter.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "获取文章",
			})
		})
		// 新增文章
		articleRouter.POST("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增文章",
			})
		})
		// 删除文章
		articleRouter.DELETE("/:cid", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除文章",
			})
		})
		// 修改文章
		articleRouter.PUT("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "修改文章",
			})
		})
	}
}
