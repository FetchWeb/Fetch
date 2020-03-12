package fetch

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

// Request type is a placeholder for *http.Request
type Request struct {
	*http.Request
	DB *gorm.DB
}
