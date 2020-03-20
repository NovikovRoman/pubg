package pubg

import (
	"log"
	"os"
	"time"
)

// For testing
var (
	cTest *Client
	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	// Не касается Matches() и Status().
	pause = time.Microsecond
)

func init() {
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		log.Fatal("Set the environment variable APIKEY before retrying.")
	}

	cTest = NewClient(os.Getenv("APIKEY"), nil)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("PAUSE") != "" {
		pause = time.Second * 8
	}
}
