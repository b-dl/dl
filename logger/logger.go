package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/kr/pretty"
	"github.com/sirupsen/logrus"
)

var (
	LogBlue = color.New(color.FgBlue)
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

func Request(r *http.Request) {
	LogBlue.Printf("URL:     ")
	fmt.Printf("%s\n", r.URL.String())
	LogBlue.Printf("Method:  ")
	fmt.Printf("%s\n", r.Method)
	LogBlue.Printf("Headers: ")
	pretty.Printf("%# v\n", r.Header)
}
