package logging

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/jubelio/go-logging/getenv"
)

var (
	apiKey   string
	username string
	password string
	host     string
	indice   string
)

func insertLogs(body LogBody) (err error) {
	host = getenv.GetEnvString("LOGGING_HOST", "https://es2.jubelio.com")
	indice = getenv.GetEnvString("LOGGING_INDICE", "logs-app-default")

	apiKey = getenv.GetEnvString("LOGGING_APIKEY", "")
	username = getenv.GetEnvString("LOGGING_USERNAME", "")
	password = getenv.GetEnvString("LOGGING_PASSWORD", "")

	if apiKey != "" {
		return insertLogsWithApiKey(body)
	}

	return insertLogsWithUserPassword(body)

}

func insertLogsWithApiKey(body LogBody) (err error) {
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		// SetResult(result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "ApiKey "+apiKey).
		// SetBasicAuth(
		// 	viper.GetString("logging.es_username"),
		// 	viper.GetString("logging.es_password")).
		SetBody(body).
		Post(fmt.Sprint("%s/%s/_doc/", host, indice))
	if err != nil {
		return
	}
	if resp.StatusCode() > 399 {
		fmt.Println(string(resp.Body()))
		err = fmt.Errorf("code: %d status: %s", resp.StatusCode(), resp.Status())
		return
	}
	return
}

func insertLogsWithUserPassword(body LogBody) (err error) {
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		// SetResult(result).
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(
			username,
			password,
		).
		SetBody(body).
		Post(fmt.Sprint("%s/%s/_doc/", host, indice))
	if err != nil {
		return
	}
	if resp.StatusCode() > 399 {
		fmt.Println(string(resp.Body()))
		err = fmt.Errorf("code: %d status: %s", resp.StatusCode(), resp.Status())
		return
	}
	return
}
