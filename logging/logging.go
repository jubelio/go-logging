package logging

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jubelio/go-logging/getenv"
)

var (
	level       string
	active      bool
	stdout      bool
	serviceName string
)

var vSeverity = [...]string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}

func Log(severity, message string, extraInfo interface{}) {
	level = getenv.GetEnvString("LOGGING_LEVEL", "INFO")
	active, _ = getenv.GetEnvBool("LOGGING_ACTIVE", false)
	stdout, _ = getenv.GetEnvBool("LOGGING_STDOUT", true)
	serviceName = getenv.GetEnvString("LOGGING_SERVICENAME", "go-logging")

	if !sContains(vSeverity[:], severity) {
		severity = "INFO"
	}

	if stdout && logLeveled(severity) {
		log.Printf("[%s] %s", severity, message)
		if extraInfo != nil {
			spew.Dump(fmt.Sprintf("[%s-extra] %s", severity, time.Now().Format("2006/01/02 15:04:05")), extraInfo)
		}
	}

	if !active || !logLeveled(severity) {
		return
	}

	logBody := LogBody{
		Timestamp: time.Now(),
		Severity:  severity,
		Message:   message,
		Fields: LogBodyFields{
			Service: serviceName,
		},
		Extra: extraInfo,
	}

	go insertLogs(logBody)

}

func logLeveled(severity string) bool {
	clvl := level
	if !sContains(vSeverity[:], clvl) {
		clvl = "INFO"
	}
	level := sIndex(vSeverity[:], clvl)
	slevel := sIndex(vSeverity[:], severity)

	return slevel <= level
}
