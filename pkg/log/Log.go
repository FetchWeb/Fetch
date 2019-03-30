package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	fetch "github.com/FetchWeb/Fetch/pkg/core"
)

// logOptions holds the currently set LogOptions
var logOptions *Options

// logFile hold the file the logs are being written to.
var logFile *os.File

// Debug for logging any messages useful for debugging applications.
func Debug(message string) {
	if logOptions.LogDebug {
		writeLog("DEBUG", message)
	}
}

// Debugf for logging any messages useful for debugging applications.
func Debugf(message string, v ...interface{}) {
	if logOptions.LogDebug {
		writeLog("DEBUG", fmt.Sprintf(message, v...))
	}
}

// Info for logging any messages about useful information.
func Info(message string) {
	if logOptions.LogInfo {
		writeLog("INFO", message)
	}
}

// Infof for logging any messages about useful information.
func Infof(message string, v ...interface{}) {
	if logOptions.LogInfo {
		writeLog("INFO", fmt.Sprintf(message, v...))
	}
}

// Warning for logging any warning messages. Example, bad practices.
func Warning(message string) {
	if logOptions.LogWarning {
		writeLog("WARNING", message)
	}
}

// Warningf for logging any warning messages. Example, bad practices.
func Warningf(message string, v ...interface{}) {
	if logOptions.LogWarning {
		writeLog("WARNING", fmt.Sprintf(message, v...))
	}
}

// Error for logging any error messages where the application is able to recover from.
func Error(message string) {
	if logOptions.LogError {
		writeLog("ERROR", message)
	}
}

// Errorf for logging any error messages where the application is able to recover from.
func Errorf(message string, v ...interface{}) {
	if logOptions.LogError {
		writeLog("ERROR", fmt.Sprintf(message, v...))
	}
}

// Fatal for logging any fatal messages where the application is unable to recover from.
func Fatal(message string) {
	if logOptions.LogFatal {
		writeLog("FATAL", message)
	}
}

// Fatalf for logging any fatal messages where the application is unable to recover from.
func Fatalf(message string, v ...interface{}) {
	if logOptions.LogFatal {
		writeLog("FATAL", fmt.Sprintf(message, v...))
	}
}

// Startup initialises the log package.
func Startup(options *Options) error {
	if options == nil {
		options = DefaultOptions()
	}

	logOptions = options

	var err error
	logFile, err = os.OpenFile(fetch.JoinStrings(logOptions.Directory, time.Now().Format("Log_2006-01-02.json")), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	output, err := json.Marshal(&Data{
		Level:     "INFO",
		Prefix:    logOptions.Prefix,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Message:   "Initial log",
	})

	if err != nil {
		return err
	}

	logFile.Write(append([]byte("["), output...))
	return nil
}

// Shutdown shuts down the log package.
func Shutdown() {
	logFile.WriteString("]")
	logFile.Close()
}

// writeLog writes the log message to console and file depending on which configuration is checked.
func writeLog(level string, message string) {
	output, _ := json.Marshal(&Data{
		Level:     level,
		Prefix:    logOptions.Prefix,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Message:   message,
	})

	logFile.Write(append([]byte(","), output...))
}
