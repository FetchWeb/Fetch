package message

import (
	"github.com/FetchWeb/Fetch/pkg/core"
)

type Email struct {
	core.QItem
	Data        *EmailData
	Credentials *EmailCredentials
}
