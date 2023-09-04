package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAuthRoute(r *gin.Engine) {
	r.POST("api/login", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "登录",
		})

	})
}
