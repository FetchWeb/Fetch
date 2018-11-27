package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var LogDirectory string

func Info(pkg string, msg string) {
	outputLog("INFO", msg, pkg)
}

func Debug(pkg string, msg string) {
	outputLog("DEBUG", msg, pkg)
}

func Warning(pkg string, msg string) {
	outputLog("WARNING", msg, pkg)
}

func Error(pkg string, msg string) {
	outputLog("ERROR", msg, pkg)
}

func Fatal(pkg string, msg string) {
	outputLog("FATAL", msg, pkg)
}

func outputLog(lvl string, msg string, pkg string) {
	currTime := time.Now()
	timestamp := currTime.Format("2006-01-02 15:04:05")

	output, err := json.Marshal(map[string]string{
		"level":   lvl,
		"time":    timestamp,
		"package": pkg,
		"message": msg})

	if err != nil {
		log.Printf("ERROR {LOG}: Failed to Marshal JSON on Info: %v", err)
		return
	}

	logFilename := currTime.Format("2006-01-02.json")
	file, err := os.OpenFile(LogDirectory+"Log_"+logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Printf("ERROR {LOG}: Failed to open log file: %v", err)
		return
	}
	defer file.Close()

	file.WriteString(string(output) + "\n")
	fmt.Println(string(output))
}
