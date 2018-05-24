package log

// Logger provides wrapper interface for logrus
type Logger interface {
	Info(v ...interface{})
	Infof(f string, v ...interface{})
	Debug(v ...interface{})
	Debugf(f string, v ...interface{})
	Warn(v ...interface{})
	Warnf(f string, v ...interface{})
	Error(v ...interface{})
	Errorf(f string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(f string, v ...interface{})
}

func (l *logger) Info(v ...interface{}) {
	l.entry.Info(v...)
}

func (l *logger) Infoln(v ...interface{}) {
	l.entry.Infoln(v...)
}

func (l *logger) Infof(f string, v ...interface{}) {
	l.entry.Infof(f, v...)
}

func (l *logger) Debug(v ...interface{}) {
	l.entry.Debug(v...)
}

func (l *logger) Debugln(v ...interface{}) {
	l.entry.Debugln(v...)
}

func (l *logger) Debugf(f string, v ...interface{}) {
	l.entry.Debugf(f, v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.entry.Warn(v...)
}

func (l *logger) Warnln(v ...interface{}) {
	l.entry.Warnln(v...)
}

func (l *logger) Warnf(f string, v ...interface{}) {
	l.entry.Warnf(f, v...)
}

func (l *logger) Error(v ...interface{}) {
	l.entry.Error(v...)
}

func (l *logger) Errorln(v ...interface{}) {
	l.entry.Errorln(v...)
}

func (l *logger) Errorf(f string, v ...interface{}) {
	l.entry.Errorf(f, v...)
}

func (l *logger) Fatal(v ...interface{}) {
	l.entry.Fatal(v...)
}

func (l *logger) Fatalln(v ...interface{}) {
	l.entry.Fatalln(v...)
}

func (l *logger) Fatalf(f string, v ...interface{}) {
	l.entry.Fatalf(f, v...)
}
