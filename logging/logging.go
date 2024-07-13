package logging

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jubelio/go-logging/getenv"
)

var (
	level       string
	active      bool
	stdout      bool
	serviceName string
	env         string
)

var vSeverity = [...]string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}

func Log(severity, message string, extraInfo interface{}) {
	level = getenv.GetEnvString("LOGGING_LEVEL", "INFO")
	active, _ = getenv.GetEnvBool("LOGGING_ACTIVE", false)
	stdout, _ = getenv.GetEnvBool("LOGGING_STDOUT", true)
	serviceName = getenv.GetEnvString("LOGGING_SERVICENAME", "go-logging")
	env = getenv.GetEnvString("ENVIRONTMENT", "development")

	if !sContains(vSeverity[:], severity) {
		severity = "INFO"
	}

	// Create a new logger instance
	logger := logrus.New()
	logger.SetReportCaller(true)

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  "2006/01/02 15:04:05",
		ForceColors:      true,
		FullTimestamp:    true,
		ForceQuote:       true,
		QuoteEmptyFields: true,
		DisableQuote:     true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			funcString := f.Function
			s := strings.Split(f.Func.Name(), ".")
			if len(s) > 1 {
				funcString = s[len(s)-1]
			}
			return fmt.Sprintf("%s()", funcString), fmt.Sprintf("%s:%d", f.File[strings.LastIndex(f.File, "/")+1:], f.Line)
		},
	})

	if stdout && logLeveled(severity) {
		var fields logrus.Fields
		extraField := logger.WithFields(logrus.Fields{})

		if extraInfo != nil {
			jsonStr, _ := json.Marshal(extraInfo)
			extraField.Data["extraInfo"] = string(jsonStr)
			fields = extraField.Data
		}

		switch severity {
		case "DEBUG":
			logger.WithFields(fields).Debug(message)
		case "INFO":
			logger.WithFields(fields).Info(message)
		case "WARN":
			logger.WithFields(fields).Warn(message)
		case "ERROR":
			logger.WithFields(fields).Error(message)
		case "FATAL":
			logger.WithFields(fields).Fatal(message)
		case "TRACE":
			logger.WithFields(fields).Trace(message)
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

// func fileInfo(skip int) string {
// 	_, file, line, ok := runtime.Caller(skip)
// 	if !ok {
// 		file = "<???>"
// 		line = 1
// 	} else {
// 		slash := strings.LastIndex(file, "/")
// 		if slash >= 0 {
// 			file = file[slash+1:]
// 		}
// 	}
// 	return fmt.Sprintf("%s:%d", file, line)
// }
