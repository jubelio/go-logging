package logging

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func sContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func sIndex(s []string, str string) int {
	for i, v := range s {
		if v == str {
			return i
		}
	}

	return -1
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[len(file)-slash:]
		}
	}

	prefix := fmt.Sprintf("%v:%v: ", path.Base(file), line)

	return prefix
}
