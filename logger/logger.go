package logger

import (
	"fmt"
	"github.com/go-logr/logr/funcr"
	"os"
)

// Logger used by the RepoAnalyzer Library, defaults to a NOOP logger
var Logger = funcr.New(
	func(pfx, args string) { fmt.Println(pfx, args) },
	funcr.Options{
		LogCaller:    funcr.All,
		LogTimestamp: true,
		Verbosity:    1,
	})

var (
	TraceLevel = -2
	DebugLevel = -1
	InfoLevel  = 0
	WarnLevel  = 1
	ErrorLevel = 2
	FatalLevel = 3
)

func Trace(msg string, keysAndValues ...interface{}) {
	Logger.V(TraceLevel).Info(msg, keysAndValues)
}
func TraceWithErr(err error, msg string, keysAndValues ...interface{}) {
	Logger.V(TraceLevel).Error(err, msg, keysAndValues)
}

func Debug(msg string, keysAndValues ...interface{}) {
	Logger.V(DebugLevel).Info(msg, keysAndValues)
}
func DebugWithErr(err error, msg string, keysAndValues ...interface{}) {
	Logger.V(DebugLevel).Error(err, msg, keysAndValues)
}

func Info(msg string, keysAndValues ...interface{}) {
	Logger.V(InfoLevel).Info(msg, keysAndValues)
}
func InfoWithErr(err error, msg string, keysAndValues ...interface{}) {
	Logger.V(InfoLevel).Error(err, msg, keysAndValues)
}

func Warn(msg string, keysAndValues ...interface{}) {
	Logger.V(WarnLevel).Info(msg, keysAndValues)
}
func WarnWithErr(err error, msg string, keysAndValues ...interface{}) {
	Logger.V(WarnLevel).Error(err, msg, keysAndValues)
}

func Error(msg string, keysAndValues ...interface{}) {
	Logger.V(ErrorLevel).Info(msg, keysAndValues)
}
func ErrorWithErr(err error, msg string, keysAndValues ...interface{}) {
	Logger.V(ErrorLevel).Error(err, msg, keysAndValues)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	Logger.V(FatalLevel).Info(msg, keysAndValues)
	os.Exit(1)
}
func FatalWithErr(err error, msg string, keysAndValues ...interface{}) {
	Logger.V(FatalLevel).Error(err, msg, keysAndValues)
	os.Exit(1)
}
