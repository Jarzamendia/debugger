package debugger

import (
	"log"
	"os"
)

//SendDebugMessage Verify a Env Variable, if true, send  a debug message.
func SendDebugMessage(msg string) {

	isDebugEnableed := getEnv("DEBUG", "false")

	if isDebugEnableed == "true" {

		log.Println(msg)

	}

}

func getEnv(key, fallback string) string {

	value := os.Getenv(key)

	if len(value) == 0 {

		return fallback

	}

	return value

}
