package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func Init(logFile string, level string) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Panic("Unable to open log file:", err)
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	l, err := logrus.ParseLevel(level)
	if err == nil {
		logrus.SetLevel(l)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	writers := []io.Writer{
		file,
		os.Stdout,
	}

	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetOutput(fileAndStdoutWriter)
	logrus.Info("Log initialization succeeded")
}
