package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// LogDirectory is the directory that the log file will be written to.
var LogDirectory string

// LogDebug is the check whether debug messages should be logged.
var LogDebug = true

// LogInfo is the check whether info messages should be logged.
var LogInfo = true

// LogWarning is the check whether warning messages should be logged.
var LogWarning = true

// LogError is the check whether error messages should be logged.
var LogError = true

// LogFatal is the check whether fatal messages should be logged.
var LogFatal = true

// LogToConsole is the check whether log messages should be written to console.
var LogToConsole = true

// LogToFile is the check whether log messages should be written to file.
var LogToFile = true

// Debug for logging any messages useful for debugging applications.
func Debug(pkg string, msg string) error {
	if LogDebug {
		return writeLog("DEBUG", msg, pkg)
	}
	return errors.New("Debug messages aren't being logged but Debug is being called")
}

// Info for logging any messages about useful information.
func Info(pkg string, msg string) error {
	if LogInfo {
		return writeLog("INFO", msg, pkg)
	}
	return errors.New("Info messages aren't being logged but Info is being called")
}

// Warning for logging any warning messages. Example, bad practices.
func Warning(pkg string, msg string) error {
	if LogWarning {
		return writeLog("WARNING", msg, pkg)
	}
	return errors.New("Warning messages aren't being logged but Warning is being called")
}

// Error for logging any error messages where the application is able to recover from.
func Error(pkg string, msg string) error {
	if LogError {
		return writeLog("ERROR", msg, pkg)
	}
	return errors.New("Error messages aren't being logged but Error is being called")
}

// Fatal for logging any fatal messages where the application is unable to recover from.
func Fatal(pkg string, msg string) error {
	if LogFatal {
		return writeLog("FATAL", msg, pkg)
	}
	return errors.New("Fatal messages aren't being logged but Fatal is being called")
}

// writeLog writes the log message to console and file depending on which configuration is checked.
func writeLog(lvl string, msg string, pkg string) error {
	currTime := time.Now()
	timestamp := currTime.Format("2006-01-02 15:04:05")

	output, err := json.Marshal(map[string]string{
		"level":   lvl,
		"time":    timestamp,
		"package": pkg,
		"message": msg})

	if err != nil {
		log.Printf("ERROR {LOG}: Failed to Marshal JSON on Info: %v", err)
		return err
	}

	if LogToFile {
		logFilename := currTime.Format("Log_2006-01-02.json")
		file, err := os.OpenFile(LogDirectory+logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

		if err != nil {
			log.Printf("ERROR {LOG}: Failed to open log file: %v", err)
			return err
		}
		defer file.Close()

		file.WriteString(string(output) + "\n")
	}

	if LogToConsole {
		fmt.Println(string(output))
	}
	return nil
}
