package logic

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func HashSecret(secret []byte) string {
	sha256 := sha256.Sum256(secret)
	return fmt.Sprintf("%x", sha256)
}

func randomBytes() ([]byte, error) {
	c := 32
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewAuthCode() (string, error) {
	bytes, err := randomBytes()
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
