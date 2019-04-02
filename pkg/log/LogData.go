package log

// Data holds all the data stored in a log.
type Data struct {
	Level     string `json:"level"`
	Prefix    string `json:"prefix"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}
