package global

import (
	"github.com/sirupsen/logrus"
	"go-blog/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Logger *logrus.Logger
)