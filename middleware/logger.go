package middleware

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"
//	"go-blog/global"
//	"io"
//	"log"
//	"os"
//	"path"
//	"time"
//)

//func LoggerToFile(FileName string) gin.HandlerFunc {
//	FilePath := global.Config.Logger.LoggerFilePath
//	fileName := path.Join(FilePath, FileName)
//	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//	if err != nil {
//		log.Fatalf("open log file failed: %s", err)
//	}
//	mw := io.MultiWriter(src, os.Stdout)
//	logger := logrus.New()
//	logger.SetOutput(mw)
//	logger.SetLevel(logrus.DebugLevel)
//	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
//	return func(c *gin.Context) {
//		startTime := time.Now()
//		c.Next()
//		endTime := time.Now()
//		latencyTime := endTime.Sub(startTime)
//		logger.Infof("%3v | %6v| %6v", startTime, endTime, latencyTime)
//	}
//}
