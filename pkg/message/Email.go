package message

import (
	"encoding/json"

	"github.com/FetchWeb/Fetch/pkg/core"
)

type Email struct {
	core.QueueItem
	Credentials *EmailCredentials
	Data        *EmailData
}

func (e *Email) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Email) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}
