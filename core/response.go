package core

import (
	"encoding/json"
	"net/http"
)

// Response type is a placeholder for http.ResponseWriter
type Response struct {
	http.ResponseWriter
}

// ToJSON writes an interface to screen as JSON
func (response Response) ToJSON(i interface{}) error {
	jsonData, err := json.Marshal(i)
	if err != nil {
		return err
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonData)

	return nil
}
