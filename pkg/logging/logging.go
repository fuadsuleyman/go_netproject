package logging

import (
	"fmt"
	"io"
	"os"
	"path"

	// "os"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer []io.Writer
	logLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.logLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func init(){
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	// allFile, err := os.Open("logs/all.log")
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Unable to write to ", allFile)
			fmt.Println(err)
			os.Exit(1)
	}
		panic(err)
	}

	l.SetOutput(io.Discard)
	l.AddHook(&writerHook{
		Writer: []io.Writer{allFile, os.Stdout},
		logLevels: logrus.AllLevels,
	})

	// for to see all
	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}