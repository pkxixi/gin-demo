package api

import (
	"go-blog/api/SettingsAapi"
	"go-blog/service"
)

type DemoApiGroup struct {
	SettingsApi SettingsAapi.SettingsApi
}

var ApiGroupApp = new(DemoApiGroup)

var (
	UserService = service.DemoServiceGroupApp.UserServiceGroup
)
