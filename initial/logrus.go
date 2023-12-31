package initial

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-blog/global"
	"io"
	"log"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 32
	blue   = 33
	gray   = 34
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//log := global.Config.Logger
	// customized date format
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// customized log file path
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// customized output format
		_, err := fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", global.Config.Logger.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
		if err != nil {
			log.Fatalf("logger can not format: %v", err)
		}
	} else {
		_, err := fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", global.Config.Logger.Prefix, timestamp, levelColor, entry.Level, entry.Message)
		if err != nil {
			log.Fatalf("logger can not format: %v", err)
		}
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	FilePath := global.Config.Logger.LoggerFilePath
	FileName := "test.log"
	fileName := path.Join(FilePath, FileName)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("open log file failed: %s", err)
	}
	mLog := logrus.New()
	mw := io.MultiWriter(os.Stdout, src)
	mLog.SetOutput(mw)
	mLog.SetReportCaller(global.Config.Logger.ShowLine)
	mLog.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	return mLog
}

//func InitDefaultLogger() {
//	logrus.SetOutput(os.Stdout)
//	logrus.SetReportCaller(global.Config.Logger.ShowLine)
//	logrus.SetFormatter(&LogFormatter{})
//	level, err := logrus.ParseLevel(global.Config.Logger.Level)
//	if err != nil {
//		level = logrus.InfoLevel
//	}
//	logrus.SetLevel(level)
//}
