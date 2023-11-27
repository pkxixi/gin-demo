package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go-blog/config"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	Config             *config.Config
	DB                 *gorm.DB
	Redis              *redis.Client
	Logger             *logrus.Logger
	ConcurrencyControl = &singleflight.Group{}
)
