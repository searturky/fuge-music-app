package utils

import (
	"github.com/google/uuid"
)

func GetRandomName(prefix string) string {
	return prefix + "_" + uuid.New().String()[:8]
}
