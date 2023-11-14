package routers

import (
	"go-blog/api"
)

func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.routerGroup.GET("settings", settingsApi.SettingsInfoView)
}
