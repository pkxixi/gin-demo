package global

import (
	"github.com/sirupsen/logrus"
	"go-blog/config"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	Config             *config.Config
	DB                 *gorm.DB
	Logger             *logrus.Logger
	ConcurrencyControl = &singleflight.Group{}
)
