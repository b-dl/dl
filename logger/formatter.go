package logger

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel, logrus.WarnLevel:
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
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	fileVal, fileValOk := entry.Data["fileVal"]
	funcVal, funcValOk := entry.Data["funcVal"]
	if fileValOk && funcValOk {
		fmt.Fprintf(b, "\x1b[%dm[%s]\x1b[0m [%s]  %s %s %s\n", levelColor, entry.Level, timestamp, fileVal, funcVal, entry.Message)
	} else if entry.HasCaller() {
		funcVal = entry.Caller.Function
		fileVal = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		fmt.Fprintf(b, "\x1b[%dm[%s]\x1b[0m [%s]  %s %s %s\n", levelColor, entry.Level, timestamp, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "\x1b[%dm[%s]\x1b[0m [%s] %s\n", levelColor, entry.Level, timestamp, entry.Message)
	}
	return b.Bytes(), nil
}
