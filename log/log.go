package log

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	/* 	file, err := os.OpenFile("log/logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	   	if err == nil {
	   		logger.Out = file
	   	} else {
	   		logger.Info("Failed to log to file, using default stderr")
	   	} */
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
