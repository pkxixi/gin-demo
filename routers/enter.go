package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
)

type RouterGroup struct {
	routerGroup *gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	ApiRouterGroup := router.Group("api")
	RouterGroupApp := RouterGroup{ApiRouterGroup}
	RouterGroupApp.SettingsRouter()
	//router.Use(middleware.LoggerToFile("test1.log"))
	return router
}
