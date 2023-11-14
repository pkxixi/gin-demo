package SettingsAapi

import (
	"github.com/gin-gonic/gin"
	"go-blog/util/response"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	response.OK(map[string]interface{}{}, "success", c)
}
