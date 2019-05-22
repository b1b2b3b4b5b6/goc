//Package logface implements a simple library
package logface

import (
	"fmt"
	"goc/toolcom/cfgtool"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

//Level is
type Level uint32

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

//Logface is
type Logface struct {
	logInner *logrus.Logger
	level    Level
}

var cfg = cfgtool.New("conf.json")

//New is
func New(argLevel Level) *Logface {
	logger := &Logface{logInner: logrus.New(), level: argLevel}
	logger.logInner.Out = rw
	logger.logInner.Level = turnLevel(argLevel)
	logger.logInner.SetFormatter(&logrus.TextFormatter{
		//DisableTimestamp: true,
	})
	return logger
}

func turnLevel(argLevel Level) logrus.Level {
	switch argLevel {
	case PanicLevel:
		return logrus.PanicLevel

	case FatalLevel:
		return logrus.FatalLevel

	case ErrorLevel:
		return logrus.ErrorLevel

	case WarnLevel:
		return logrus.WarnLevel

	case InfoLevel:
		return logrus.InfoLevel

	case DebugLevel:
		return logrus.DebugLevel

	case TraceLevel:
		return logrus.TraceLevel
	}
	os.Exit(-1)
	return logrus.TraceLevel
}

//Logkit is
func (p *Logface) Logkit(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Trace(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Trace is
func (p *Logface) Trace(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Trace(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Debug is
func (p *Logface) Debug(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Debug(fmt.Sprintf("[%s.%d] ", file, line), fmt.Sprintf(format, args...))
}

//Info is
func (p *Logface) Info(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Info(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Print is
func (p *Logface) Print(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Print(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Warn is
func (p *Logface) Warn(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Warn(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Warning is
func (p *Logface) Warning(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Warning(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Error is
func (p *Logface) Error(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Error(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Fatal is
func (p *Logface) Fatal(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Fatal(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}

//Panic is
func (p *Logface) Panic(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	offset := strings.LastIndex(file, "/")
	file = file[offset+1:]
	p.logInner.Panic(fmt.Sprintf("[%s][%d] ", file, line), fmt.Sprintf(format, args...))
}
