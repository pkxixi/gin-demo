package initial

import (
	"go-blog/global"
	"go-blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

func Mysql() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		//global.LOG.warn
		global.Logger.Warnf("为配置host，取消连接！\n")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Logger.Fatalf("[%s] mysql连接失败。", dsn)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(models.User{})
	if err != nil {
		global.Logger.Error("register table failed")
		os.Exit(0)
	}
	global.Logger.Info("register table success")
}

//func MysqlRegisterTables() {
//	db := global.DB
//	err := db.AutoMigrate()
//	if err != nil {
//		global.Logger.Errorf("register table failed: %v", err)
//		os.Exit(0)
//	}
//	global.Logger.Info("register table success")
//}
