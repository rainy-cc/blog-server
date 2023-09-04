package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCategoryRoute(r *gin.Engine) {
	cateRouter := r.Group("api/category")
	{
		// 获取所有分类
		cateRouter.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "查询分类",
			})
		})
		// 查询单个分类
		cateRouter.GET("/:cid", func(ctx *gin.Context) {
			cid := ctx.Param("cid")
			cidInt, err := strconv.Atoi(cid)
			if err != nil {
				fmt.Println(err.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"code": 400,
					"msg":  "参数错误",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "查询单个分类",
				"id":   cidInt,
			})
		})
		// 新增分类
		cateRouter.POST("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增分类",
			})
		})
		// 删除分类
		cateRouter.DELETE("/:cid", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除分类",
			})
		})
		// 修改分类
		cateRouter.PUT("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "修改分类",
			})
		})
	}
}
