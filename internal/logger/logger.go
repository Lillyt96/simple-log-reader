package logger

import (
	"github.com/sirupsen/logrus"
)

func Default() *logrus.Logger {
	defaultLog := logrus.New()
	defaultLog.SetFormatter(&logrus.TextFormatter{})
	defaultLog.SetLevel(logrus.InfoLevel)

	return defaultLog
}
