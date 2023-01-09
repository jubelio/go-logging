package logging

import (
	"fmt"
	"runtime"
)

func Debugf(message string, extraInfo interface{}) { Log("DEBUG", message, extraInfo) }
func Infof(message string, extraInfo interface{})  { Log("INFO", message, extraInfo) }
func Warnf(message string, extraInfo interface{})  { Log("WARN", message, extraInfo) }
func Errorf(message string, extraInfo interface{}) { Log("ERROR", message, extraInfo) }
func Fatalf(message string, extraInfo interface{}) { Log("FATAL", message, extraInfo) }
func Tracef(message string, extraInfo interface{}) {
	pc, filename, line, _ := runtime.Caller(1)

	x := fmt.Sprintf("[%s[%s:%d]] ", runtime.FuncForPC(pc).Name(), filename, line)
	Log("TRACE", x+message, extraInfo)
}

func Debug(message string) { Log("DEBUG", message, nil) }
func Info(message string)  { Log("INFO", message, nil) }
func Warn(message string)  { Log("WARN", message, nil) }
func Error(message string) { Log("ERROR", message, nil) }
func Fatal(message string) { Log("FATAL", message, nil) }
func Trace(message string) {
	pc, filename, line, _ := runtime.Caller(1)

	x := fmt.Sprintf("[%s[%s:%d]] ", runtime.FuncForPC(pc).Name(), filename, line)
	Log("TRACE", x+message, nil)
}
