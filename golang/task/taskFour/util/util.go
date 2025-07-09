package util

import (
	"crypto/rand"
)

func GenerateSecureKey(lengthInBytes int) ([]byte, error) {
	key := make([]byte, lengthInBytes)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
