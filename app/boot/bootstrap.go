package boot

import (
	"github.com/gin-gonic/gin"
	"gogo/app/routes"
)


// 启动入口
func Bootstrap(c *gin.Engine) {
	SetRoute(c)
}

// 路由入口
func SetRoute(c *gin.Engine) {
	routes.Route(c)
}