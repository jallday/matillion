package utils

import (
	"github.com/google/uuid"
)

func ValidId(id string) bool {
	return len(id) == 36
}

func NewID() string {
	return uuid.New().String()
}
