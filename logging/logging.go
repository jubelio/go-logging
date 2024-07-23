package logging

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jubelio/go-logging/getenv"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	logger      = logrus.New()
	level       string
	active      bool
	serviceName string
	vSeverity   = [...]string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

func init() {
	level = getenv.GetEnvString("LOGGING_LEVEL", "INFO")
	active, _ = getenv.GetEnvBool("LOGGING_ACTIVE", false)
	serviceName = getenv.GetEnvString("LOGGING_SERVICENAME", "go-logging")

	logger.SetFormatter(
		&prefixed.TextFormatter{
			TimestampFormat:  "2006-01-02 15:04:05",
			ForceColors:      true,
			FullTimestamp:    true,
			QuoteEmptyFields: true,
			ForceFormatting:  true,
			DisableUppercase: true,
		})

	logger.SetReportCaller(true)
	logger.SetLevel(logrus.InfoLevel)
}

func sendLog(severity, message string, extraInfo interface{}) {
	if !sContains(vSeverity[:], severity) {
		log.Println("severity", severity)

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

	// send logs to ElasticSearch
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

func Infof(message string, extraInfo interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		if extraInfo != nil && extraInfo != "" {
			jsonStr, _ := json.Marshal(extraInfo)
			entity.Data["data"] = string(jsonStr)
		}
		entity.Info(message)
		go sendLog("INFO", message, extraInfo)
	}
}

func Warnf(message string, extraInfo interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		if extraInfo != nil && extraInfo != "" {
			jsonStr, _ := json.Marshal(extraInfo)
			entity.Data["data"] = string(jsonStr)
		}

		entity.Warn(message)
		go sendLog("WARN", message, extraInfo)
	}
}

func Errorf(message string, extraInfo interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		if extraInfo != nil && extraInfo != "" {
			if err, ok := extraInfo.(error); ok {
				entity.Data["reason"] = err.Error()
			} else {
				jsonStr, _ := json.Marshal(extraInfo)
				entity.Data["data"] = string(jsonStr)
			}
		}
		entity.Error(message)
		go sendLog("ERROR", message, extraInfo)
	}
}

func Fatalf(message string, extraInfo interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		if extraInfo != nil && extraInfo != "" {
			if err, ok := extraInfo.(error); ok {
				entity.Data["reason"] = err.Error()
			} else {
				jsonStr, _ := json.Marshal(extraInfo)
				entity.Data["data"] = string(jsonStr)
			}
		}

		entity.Fatal(message)
		go sendLog("FATAL", message, extraInfo)
	}
}

func Tracef(message string, extraInfo interface{}) {
	if logger.Level >= logrus.TraceLevel {
		entity := logger.WithFields(logrus.Fields{})
		if extraInfo != nil && extraInfo != "" {
			jsonStr, _ := json.Marshal(extraInfo)
			entity.Data["data"] = string(jsonStr)
		}

		entity.Trace(message)
		go sendLog("TRACE", message, extraInfo)
	}
}

func Debugf(message string, extraInfo interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		if extraInfo != nil && extraInfo != "" {
			jsonStr, _ := json.Marshal(extraInfo)
			entity.Data["data"] = string(jsonStr)
		}

		entity.Debug(message)
		go sendLog("DEBUG", message, extraInfo)
	}
}

func Trace(message string) {
	if logger.Level >= logrus.TraceLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		entity.Trace(message)
		go sendLog("TRACE", message, nil)
	}
}

func Debug(message string) {
	if logger.Level >= logrus.DebugLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		entity.Debug(message)
		go sendLog("DEBUG", message, nil)
	}
}

func Info(message string) {
	if logger.Level >= logrus.InfoLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		entity.Infof(message)
		go sendLog("INFO", message, nil)
	}
}

func Warn(message string) {
	if logger.Level >= logrus.WarnLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		entity.Warn(message)
		go sendLog("WARN", message, nil)
	}
}

func Error(message string) {
	if logger.Level >= logrus.ErrorLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		entity.Error(message)
		go sendLog("ERROR", message, nil)
	}
}

func Fatal(message string) {
	if logger.Level >= logrus.FatalLevel {
		entity := logger.WithFields(logrus.Fields{
			"file": fileInfo(2),
		})
		entity.Fatal(message)
		go sendLog("FATAL", message, nil)
	}
}

// WithFields returns a new entry with the given fields.
// It is a shortcut for `WithFields(Fields(fields))`.
// See `WithFields` for more details.
func WithFields(fields Fields) *logrus.Entry {
	return logger.WithFields(logrus.Fields(fields))
}
