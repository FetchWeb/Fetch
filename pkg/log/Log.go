package log

import (
	"encoding/json"
	"log"
	"os"
	"time"

	fetch "github.com/FetchWeb/Fetch/pkg/core"
)

// logOptions holds the currently set LogOptions
var logOptions *Options

// logFile hold the file the logs are being written to.
var logFile *os.File

// Debug for logging any messages useful for debugging applications.
func Debug(message string) error {
	if logOptions.LogDebug {
		return writeLog("DEBUG", message)
	}
	return nil
}

// Info for logging any messages about useful information.
func Info(message string) error {
	if logOptions.LogInfo {
		return writeLog("INFO", message)
	}
	return nil
}

// Warning for logging any warning messages. Example, bad practices.
func Warning(message string) error {
	if logOptions.LogWarning {
		return writeLog("WARNING", message)
	}
	return nil
}

// Error for logging any error messages where the application is able to recover from.
func Error(message string) error {
	if logOptions.LogError {
		return writeLog("ERROR", message)
	}
	return nil
}

// Fatal for logging any fatal messages where the application is unable to recover from.
func Fatal(message string) error {
	if logOptions.LogFatal {
		return writeLog("FATAL", message)
	}
	return nil
}

// Startup initialises the log package.
func Startup(options *Options) {
	if options == nil {
		options = DefaultOptions()
	}

	logOptions = options

	var err error
	logFile, err = os.OpenFile(fetch.JoinStrings(logOptions.Directory, time.Now().Format("Log_2006-01-02.json")), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
	}

	output, err := json.Marshal(&Data{
		Level:     "INFO",
		Prefix:    logOptions.Prefix,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Message:   "Initial log",
	})

	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	}

	logFile.Write(append([]byte("["), output...))
}

// Shutdown shuts down the log package.
func Shutdown() {
	logFile.WriteString("]")
	logFile.Close()
}

// writeLog writes the log message to console and file depending on which configuration is checked.
func writeLog(level string, message string) error {
	output, err := json.Marshal(&Data{
		Level:     level,
		Prefix:    logOptions.Prefix,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Message:   message,
	})

	if err != nil {
		return err
	}

	logFile.Write(append([]byte(","), output...))

	return nil
}
