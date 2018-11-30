package log_test

import (
	"os"
	"testing"
	"time"

	"github.com/FetchWeb/Fetch/pkg/log"
)

const logDirectory = ""

func TestLogging(t *testing.T) {
	log.LogDirectory = logDirectory

	err := log.Info("LOG_TEST", "This is an a test info message")
	if err != nil {
		t.Errorf("Failed to log Info message: %v", err)
	}

	log.Debug("LOG_TEST", "This is a test debug message")
	if err != nil {
		t.Errorf("Failed to log Debug message: %v", err)
	}

	log.Warning("LOG_TEST", "This is a test warning message")
	if err != nil {
		t.Errorf("Failed to log Warning message: %v", err)
	}

	log.Error("LOG_TEST", "This is an test error message")
	if err != nil {
		t.Errorf("Failed to log Error message: %v", err)
	}

	log.Fatal("LOG_TEST", "This is a test fatal message")
	if err != nil {
		t.Errorf("Failed to log Fatal message: %v", err)
	}

	if err := os.Remove(time.Now().Format("Log_2006-01-02.json")); err != nil {
		t.Logf("Failed to delete log file after test: %v", err)
	}
}

func TestLoggingToConsole(t *testing.T) {
	log.LogDirectory = logDirectory
	log.LogToFile = false

	err := log.Info("LOG_TEST", "This is an a test info message")
	if err != nil {
		t.Errorf("Failed to log Info message: %v", err)
	}

	log.Debug("LOG_TEST", "This is a test debug message")
	if err != nil {
		t.Errorf("Failed to log Debug message: %v", err)
	}

	log.Warning("LOG_TEST", "This is a test warning message")
	if err != nil {
		t.Errorf("Failed to log Warning message: %v", err)
	}

	log.Error("LOG_TEST", "This is an test error message")
	if err != nil {
		t.Errorf("Failed to log Error message: %v", err)
	}

	log.Fatal("LOG_TEST", "This is a test fatal message")
	if err != nil {
		t.Errorf("Failed to log Fatal message: %v", err)
	}
}

func TestLoggingToFile(t *testing.T) {
	log.LogDirectory = logDirectory
	log.LogToConsole = false

	err := log.Info("LOG_TEST", "This is an a test info message")
	if err != nil {
		t.Errorf("Failed to log Info message: %v", err)
	}

	log.Debug("LOG_TEST", "This is a test debug message")
	if err != nil {
		t.Errorf("Failed to log Debug message: %v", err)
	}

	log.Warning("LOG_TEST", "This is a test warning message")
	if err != nil {
		t.Errorf("Failed to log Warning message: %v", err)
	}

	log.Error("LOG_TEST", "This is an test error message")
	if err != nil {
		t.Errorf("Failed to log Error message: %v", err)
	}

	log.Fatal("LOG_TEST", "This is a test fatal message")
	if err != nil {
		t.Errorf("Failed to log Fatal message: %v", err)
	}

	if err := os.Remove(time.Now().Format("Log_2006-01-02.json")); err != nil {
		t.Logf("Failed to delete log file after test: %v", err)
	}
}

func BenchmarkLogging(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.LogDirectory = logDirectory

		err := log.Info("LOG_TEST", "This is an a test info message")
		if err != nil {
			b.Errorf("Failed to log Info message: %v", err)
		}

		log.Debug("LOG_TEST", "This is a test debug message")
		if err != nil {
			b.Errorf("Failed to log Debug message: %v", err)
		}

		log.Warning("LOG_TEST", "This is a test warning message")
		if err != nil {
			b.Errorf("Failed to log Warning message: %v", err)
		}

		log.Error("LOG_TEST", "This is an test error message")
		if err != nil {
			b.Errorf("Failed to log Error message: %v", err)
		}

		log.Fatal("LOG_TEST", "This is a test fatal message")
		if err != nil {
			b.Errorf("Failed to log Fatal message: %v", err)
		}
	}

	if err := os.Remove(time.Now().Format("Log_2006-01-02.json")); err != nil {
		b.Logf("Failed to delete log file after test: %v", err)
	}
}

func BenchmarkLoggingToConsole(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.LogDirectory = logDirectory
		log.LogToFile = false

		err := log.Info("LOG_TEST", "This is an a test info message")
		if err != nil {
			b.Errorf("Failed to log Info message: %v", err)
		}

		log.Debug("LOG_TEST", "This is a test debug message")
		if err != nil {
			b.Errorf("Failed to log Debug message: %v", err)
		}

		log.Warning("LOG_TEST", "This is a test warning message")
		if err != nil {
			b.Errorf("Failed to log Warning message: %v", err)
		}

		log.Error("LOG_TEST", "This is an test error message")
		if err != nil {
			b.Errorf("Failed to log Error message: %v", err)
		}

		log.Fatal("LOG_TEST", "This is a test fatal message")
		if err != nil {
			b.Errorf("Failed to log Fatal message: %v", err)
		}
	}
}

func BenchmarkLoggingToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.LogDirectory = logDirectory
		log.LogToConsole = false

		err := log.Info("LOG_TEST", "This is an a test info message")
		if err != nil {
			b.Errorf("Failed to log Info message: %v", err)
		}

		log.Debug("LOG_TEST", "This is a test debug message")
		if err != nil {
			b.Errorf("Failed to log Debug message: %v", err)
		}

		log.Warning("LOG_TEST", "This is a test warning message")
		if err != nil {
			b.Errorf("Failed to log Warning message: %v", err)
		}

		log.Error("LOG_TEST", "This is an test error message")
		if err != nil {
			b.Errorf("Failed to log Error message: %v", err)
		}

		log.Fatal("LOG_TEST", "This is a test fatal message")
		if err != nil {
			b.Errorf("Failed to log Fatal message: %v", err)
		}
	}

	if err := os.Remove(time.Now().Format("Log_2006-01-02.json")); err != nil {
		b.Logf("Failed to delete log file after test: %v", err)
	}
}
