package logging

import (
	"fmt"
	"runtime"
)

// Debugf - Debug severity with extrainfo
func Debugf(message string, extraInfo interface{}) { Log("DEBUG", message, extraInfo) }

// Infof - Info severity with extrainfo
func Infof(message string, extraInfo interface{}) { Log("INFO", message, extraInfo) }

// Warnf - Warn severity with extrainfo
func Warnf(message string, extraInfo interface{}) { Log("WARN", message, extraInfo) }

// Errorf - Error severity with extrainfo
func Errorf(message string, extraInfo interface{}) { Log("ERROR", message, extraInfo) }

// Fatalf - Fatal severity with extrainfo
func Fatalf(message string, extraInfo interface{}) { Log("FATAL", message, extraInfo) }

// Tracef - Trace severity with extrainfo
func Tracef(message string, extraInfo interface{}) {
	pc, filename, line, _ := runtime.Caller(1)

	x := fmt.Sprintf("[%s[%s:%d]] ", runtime.FuncForPC(pc).Name(), filename, line)
	Log("TRACE", x+message, extraInfo)
}

// Debug - Debug severity logging
func Debug(message string) { Log("DEBUG", message, nil) }

// Info - Info severity logging
func Info(message string) { Log("INFO", message, nil) }

// Warn - Warn severity logging
func Warn(message string) { Log("WARN", message, nil) }

// Error - Error severity logging
func Error(message string) { Log("ERROR", message, nil) }

// Fatal - Fatal severity logging
func Fatal(message string) { Log("FATAL", message, nil) }

// Trace - Trace severity logging
func Trace(message string) {

	pc, filename, line, _ := runtime.Caller(1)

	x := fmt.Sprintf("[%s[%s:%d]] ", runtime.FuncForPC(pc).Name(), filename, line)
	Log("TRACE", x+message, nil)
}
