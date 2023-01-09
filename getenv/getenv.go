package getenv

import (
	"os"
	"strconv"
	"strings"
)

// GetEnvString gets the environment variable for a key and if that env-var hasn't been set it returns the default value
func GetEnvString(key string, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultVal
	}
	return value
}

// GetEnvString gets the environment variable for a key and if that env-var hasn't been set it returns the default value
func GetEnvArrString(key, separator string, defaultVal []string) []string {
	value := strings.Split(os.Getenv(key), separator)
	if len(value) == 0 {
		value = defaultVal
	}
	return value
}

// GetEnvBool gets the environment variable for a key and if that env-var hasn't been set it returns the default value
func GetEnvBool(key string, defaultVal bool) (bool, error) {
	envvalue := os.Getenv(key)
	if len(envvalue) == 0 {
		value := defaultVal
		return value, nil
	}

	value, err := strconv.ParseBool(envvalue)
	return value, err
}

// GetEnvInt gets the environment variable for a key and if that env-var hasn't been set it returns the default value. This function is equivalent to ParseInt(s, 10, 0) to convert env-vars to type int
func GetEnvInt(key string, defaultVal int) (int, error) {
	envvalue := os.Getenv(key)
	if len(envvalue) == 0 {
		value := defaultVal
		return value, nil
	}

	value, err := strconv.Atoi(envvalue)
	return value, err
}

// GetEnvFloat gets the environment variable for a key and if that env-var hasn't been set it returns the default value. This function uses bitSize of 64 to convert string to float64.
func GetEnvFloat(key string, defaultVal float64) (float64, error) {
	envvalue := os.Getenv(key)
	if len(envvalue) == 0 {
		value := defaultVal
		return value, nil
	}

	value, err := strconv.ParseFloat(envvalue, 64)
	return value, err
}
