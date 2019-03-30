package log

// Options @TODO
type Options struct {
	LogDebug   bool
	LogInfo    bool
	LogWarning bool
	LogError   bool
	LogFatal   bool
	Prefix     string
	Directory  string
}

// DefaultOptions @TODO
func DefaultOptions() *Options {
	return &Options{
		LogDebug:   true,
		LogInfo:    true,
		LogWarning: true,
		LogError:   true,
		LogFatal:   true,
		Prefix:     "",
		Directory:  "",
	}
}
