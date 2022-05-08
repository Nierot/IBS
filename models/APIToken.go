package models

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"gorm.io/gorm"
)

type APIToken struct {
	gorm.Model
	Token    string    `json:"Token"`
	Device   string    `json:"Device"`
	LastUsed time.Time `json:"LastUsed"`
}

func GenerateToken(device string) *APIToken {
	token, err := generateRandomStringURLSafe(64)

	// If this fails then shit is really fucked
	if err != nil {
		panic(err)
	}

	return &APIToken{
		Token:    token,
		Device:   device,
		LastUsed: time.Unix(1, 0),
	}
}

// Credits to https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomStringURLSafe(n int) (string, error) {
	b, err := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
