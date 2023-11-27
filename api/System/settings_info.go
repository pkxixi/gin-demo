package System

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/models/response"
)

type SettingsApi struct{}

func (s *SettingsApi) SettingsInfoView(c *gin.Context) {
	global.Logger.Info("this is the first log test")
	response.OK(map[string]interface{}{}, "success", c)
}
