package SettingsAapi

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/util/response"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	global.Logger.Info("this is the first log test")
	response.OK(map[string]interface{}{}, "success", c)
}
