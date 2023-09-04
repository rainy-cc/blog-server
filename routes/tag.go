package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTagRoute(r *gin.Engine) {
	tagRouter := r.Group("api/tag")
	{
		// 获取所有标签
		tagRouter.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "查询所有标签",
			})
		})
		//// 获取单个标签
		//tagRouter.GET("/:tagId", func(ctx *gin.Context) {
		//	ctx.JSON(http.StatusOK, gin.H{
		//		"code": 0,
		//		"msg":  "查询所有标签",
		//	})
		//})
		// 新增标签
		tagRouter.POST("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增标签",
			})
		})
		// 删除标签
		tagRouter.DELETE("/:tagId", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除标签",
			})
		})
		// 修改标签
		tagRouter.PUT("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "修改标签",
			})
		})
	}
}
