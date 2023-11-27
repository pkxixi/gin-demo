package system

import (
	"github.com/gin-gonic/gin"
	"go-blog/api"
)

type RouterSystem struct{}

func (r *RouterSystem) SettingsRouter(Router *gin.RouterGroup) {
	//settingsApi := api.GroupApi.ApiGroup
	//r.routerGroup.GET("settings", settingsApi.SettingsInfoView)
	//systemRouterWithoutRecord := Router.Group("system")
	systemApi := api.GroupApi.SystemApiGroup.SettingsApi
	{
		//systemRouterWithoutRecord.GET("settings", systemApi.SettingsInfoView)
		Router.GET("settings", systemApi.SettingsInfoView)
	}
}
