package log_test

import (
	"os"
	"testing"
	"time"

	fetch "github.com/FetchWeb/Fetch/pkg/core"
	log "github.com/FetchWeb/Fetch/pkg/log"
)

func TestLogging(t *testing.T) {
	log.Startup(&log.Options{
		LogDebug:   true,
		LogInfo:    true,
		LogWarning: true,
		LogError:   true,
		LogFatal:   true,
		Prefix:     "TEST",
		Directory:  "logs/",
	})

	if err := log.Info("This is a test info message"); err != nil {
		t.Errorf("Failed to log Info message: %v", err)
	}

	if err := log.Debug("This is a test debug message"); err != nil {
		t.Errorf("Failed to log Debug message: %v", err)
	}

	if err := log.Warning("This is a test warning message"); err != nil {
		t.Errorf("Failed to log Warning message: %v", err)
	}

	if err := log.Error("This is a test error message"); err != nil {
		t.Errorf("Failed to log Error message: %v", err)
	}

	if err := log.Fatal("This is a test fatal message"); err != nil {
		t.Errorf("Failed to log Fatal message: %v", err)
	}

	log.Shutdown()

	if err := os.Remove(fetch.JoinStrings("logs/", time.Now().Format("Log_2006-01-02.json"))); err != nil {
		t.Logf("Failed to delete log file after test: %v", err)
	}
}

func BenchmarkLogging(b *testing.B) {
	log.Startup(&log.Options{
		LogDebug:   true,
		LogInfo:    true,
		LogWarning: true,
		LogError:   true,
		LogFatal:   true,
		Prefix:     "TEST",
		Directory:  "logs/",
	})

	for i := 0; i < b.N; i++ {
		if err := log.Info("This is a test info message"); err != nil {
			b.Errorf("Failed to log Info message: %v", err)
		}

		if err := log.Debug("This is a test debug message"); err != nil {
			b.Errorf("Failed to log Debug message: %v", err)
		}

		if err := log.Warning("This is a test warning message"); err != nil {
			b.Errorf("Failed to log Warning message: %v", err)
		}

		if err := log.Error("This is a test error message"); err != nil {
			b.Errorf("Failed to log Error message: %v", err)
		}

		if err := log.Fatal("This is a test fatal message"); err != nil {
			b.Errorf("Failed to log Fatal message: %v", err)
		}

		log.Shutdown()
	}

	if err := os.Remove(fetch.JoinStrings("logs/", time.Now().Format("Log_2006-01-02.json"))); err != nil {
		b.Logf("Failed to delete log file after test: %v", err)
	}
}
