package log

// Options holds all the possible options that can be set when loggin.
type Options struct {
	LogDebug   bool
	LogInfo    bool
	LogWarning bool
	LogError   bool
	LogFatal   bool
	Prefix     string
	Directory  string
}

// DefaultOptions are the default options that can be set when logging.
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
