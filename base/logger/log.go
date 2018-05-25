package logx

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type logger struct {
	logrusLogger *logrus.Logger
}

// New to init logger
func New() Logger {
	dateNowInYMD := time.Now().Format("20060102")
	path := "./eventlog/" + dateNowInYMD

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
	eWriter, _ := rotatelogs.New(
		path+"/error.log",
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	iWriter, _ := rotatelogs.New(
		path+"/info.log",
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)

	log := &logger{
		logrusLogger: logrus.New(),
	}

	log.logrusLogger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  iWriter,
			logrus.ErrorLevel: eWriter,
		},
		&logrus.JSONFormatter{},
	))

	return log
}

// Logger provides wrapper interface for logrus
type Logger interface {
	Info(args ...interface{})
	// Infof(f string, args ...interface{})
	// Debug(args ...interface{})
	// Debugf(f string, args ...interface{})
	// Warn(args ...interface{})
	// Warnf(f string, args ...interface{})
	Error(args ...interface{})
	// Errorf(f string, args ...interface{})
	// Fatal(args ...interface{})
	// Fatalf(f string, args ...interface{})
}

func (l *logger) Info(args ...interface{}) {
	l.logrusLogger.WithFields(logrus.Fields{
		"source": fileInfo(2),
	}).Info(args...)
}

// func(l *logger)Infoln(args ...interface{}) {
// 	l.logger.Infoln(args...)
// }

// func(l *logger)Infof(f string, args ...interface{}) {
// 	l.logger.Infof(f, args...)
// }

// func(l *logger)Debug(args ...interface{}) {
// 	l.logger.Debug(args...)
// }

// func(l *logger)Debugln(args ...interface{}) {
// 	l.logger.Debugln(args...)
// }

// func(l *logger)Debugf(f string, args ...interface{}) {
// 	l.logger.Debugf(f, args...)
// }

// func(l *logger)Warn(args ...interface{}) {
// 	l.logger.Warn(args...)
// }

// func(l *logger)Warnln(args ...interface{}) {
// 	l.logger.Warnln(args...)
// }

// func(l *logger)Warnf(f string, args ...interface{}) {
// 	l.logger.Warnf(f, args...)
// }

func (l *logger) Error(args ...interface{}) {
	l.logrusLogger.WithFields(logrus.Fields{
		"source": fileInfo(2),
	}).Error(args...)
}

// func(l *logger)Errorln(args ...interface{}) {
// 	l.logger.Errorln(args...)
// }

// func(l *logger)Errorf(f string, args ...interface{}) {
// 	l.logger.Errorf(f, args...)
// }

// func(l *logger)Fatal(args ...interface{}) {
// 	l.logger.Fatal(args...)
// }

// func(l *logger)Fatalln(args ...interface{}) {
// 	l.logger.Fatalln(args...)
// }

// func(l *logger)Fatalf(f string, args ...interface{}) {
// 	l.logger.Fatalf(f, args...)
// }

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
