package api

import "go-blog/api/SettingsAapi"

type ApiGroup struct {
	SettingsApi SettingsAapi.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
