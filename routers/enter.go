package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/routers/system"
	v1 "go-blog/routers/v1"
)

type RouterGroup struct {
	V1          v1.RouterGroupV1
	SystemGroup system.RouterGroupSystem
}

var RG = new(RouterGroup)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	systemRouter := RG.SystemGroup.RouterSystem
	v1Router := RG.V1.RouterUser
	settingsRouterGroup := router.Group("system")
	usersRouterGroup := router.Group("user")
	{
		systemRouter.SettingsRouter(settingsRouterGroup)
	}
	{
		v1Router.InitUserRouter(usersRouterGroup)
	}
	//RouterGroupApp := RouterGroup{ApiRouterGroup}
	//RouterGroupApp.SettingsRouter()
	//router.Use(middleware.LoggerToFile("test1.log"))
	return router
}
