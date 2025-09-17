package gologger

import "log"

func LogInfo(message string) {
	log.Printf("INFO: - %v ", message)
}

func LogWarning(message string) {
	log.Printf("WARN - %v ", message)
}

func LogError(message string) {
	log.Printf("ERROR - %v ", message)
}

func LogDebug(message string) {
	log.Printf("DEBUG - %v ", message)
}
