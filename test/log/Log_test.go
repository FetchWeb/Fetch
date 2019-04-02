package log_test

import (
	"errors"
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

	log.Info("This is a test info message")
	log.Debug("This is a test debug message")
	log.Warning("This is a test warning message")
	log.Error("This is a test error message")
	log.Fatal("This is a test fatal message")
	
	log.Infof("This is a test info %v", errors.New("message"))
	log.Debugf("This is a test debug %v", errors.New("message"))
	log.Warningf("This is a test warning %v", errors.New("message"))
	log.Errorf("This is a test error %v", errors.New("message"))
	log.Fatalf("This is a test fatal %v", errors.New("message"))

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
		log.Info("This is a test info message")
		log.Debug("This is a test debug message")
		log.Warning("This is a test warning message")
		log.Error("This is a test error message")
		log.Fatal("This is a test fatal message")

		log.Infof("This is a test info %v", errors.New("message"))
		log.Debugf("This is a test debug %v", errors.New("message"))
		log.Warningf("This is a test warning %v", errors.New("message"))
		log.Errorf("This is a test error %v", errors.New("message"))
		log.Fatalf("This is a test fatal %v", errors.New("message"))
	}
	
	log.Shutdown()

	if err := os.Remove(fetch.JoinStrings("logs/", time.Now().Format("Log_2006-01-02.json"))); err != nil {
		b.Logf("Failed to delete log file after test: %v", err)
	}
}
