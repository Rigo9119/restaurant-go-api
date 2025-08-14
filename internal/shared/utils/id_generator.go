package utils

import (
	"fmt"
	"time"
)

// GenerateRandomID creates a unique ID with a prefix
// Example: GenerateRandomID("user") -> "user_1645123456789"
func GenerateRandomID(prefix string) string {
	return fmt.Sprintf("%s_%d", prefix, time.Now().UnixNano())
}