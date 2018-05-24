package log

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Level mirror from logrus
type Level int

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type logger struct {
	entry *logrus.Entry
}

// New returns the logging entry with contextName
func New() Logger {
	logger := &logrus.Entry{}
	logger = logrus.WithFields(logrus.Fields{
		"source": fileInfo(2),
	})

	formatter := new(LogFormatter)
	formatter.TimestampFormat = "01-Jan-2006 15:04:05"
	formatter.LevelDesc = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"}
	logrus.SetFormatter(formatter)

	return logger
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// SetLogLevel sets the global log level
func SetLogLevel(level Level) {
	logrus.SetLevel(logrus.Level(level))
}

// SetLogOutput sets the global logging output destination
func SetLogOutput(w io.Writer) {
	logrus.SetOutput(w)
}
