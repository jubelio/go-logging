package logging

import (
	"fmt"
	"runtime"
)

func DEBUG(message string, extraInfo interface{}) { Log("DEBUG", message, extraInfo) }
func INFO(message string, extraInfo interface{})  { Log("INFO", message, extraInfo) }
func WARN(message string, extraInfo interface{})  { Log("WARN", message, extraInfo) }
func ERROR(message string, extraInfo interface{}) { Log("ERROR", message, extraInfo) }
func FATAL(message string, extraInfo interface{}) { Log("FATAL", message, extraInfo) }
func TRACE(message string, extraInfo interface{}) {
	pc, filename, line, _ := runtime.Caller(1)

	x := fmt.Sprintf("[%s[%s:%d]] ", runtime.FuncForPC(pc).Name(), filename, line)
	Log("TRACE", x+message, extraInfo)
}
