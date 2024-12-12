package utils

import (
	"log"
	"os"
)

func InitLogger() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Log(v ...interface{}) {
	log.Println(v...)
}
