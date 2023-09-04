package initialize

import (
	"blog/routes"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	routes.LoadAllRoutes(r)
	return r
}
