package strutil

import (
	"github.com/google/uuid"
)

// UUID returns the string form of uuid, 12345678-1234-1234-1234-1234567890x .
func UUID() string {
	return uuid.New().String()
}
