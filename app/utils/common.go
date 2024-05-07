package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fuge/app/core"

	"github.com/google/uuid"
)

func GetRandomName(prefix string) string {
	return prefix + "_" + uuid.New().String()[:8]
}

func GetRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}

func GenRandomNickname() string {
	prefix := core.GetConf().DefaultUserPrefix
	randomName := GetRandomName(prefix)
	return randomName
}
