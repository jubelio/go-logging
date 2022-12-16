package logging

import (
	"log"
	"time"

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

	err := insertLogs(logBody)
	if err != nil {
		log.Printf("[insertlog-%s] %s", "ERROR", err.Error())
	}
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
