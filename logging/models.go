package logging

import "time"

type LogBody struct {
	Timestamp time.Time     `json:"@timestamp"`
	Severity  string        `json:"severity,omitempty"`
	Fields    LogBodyFields `json:"fields,omitempty"`
	Extra     interface{}   `json:"extra,omitempty"`
	Message   string        `json:"message,omitempty"`
}

type LogBodyFields struct {
	Service string `json:"service,omitempty"`
}
