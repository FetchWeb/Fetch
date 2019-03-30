package log

// Data @TODO
type Data struct {
	Level     string `json:"level"`
	Prefix    string `json:"prefix"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}
