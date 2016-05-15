package token

import (
	"crypto/rand"
	"encoding/base64"
)

// generate generates a base64 token from a byte array of given length
func generate(length int) (string, error) {
	rb := make([]byte, length)

	_, err := rand.Read(rb)

	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(rb), nil
}

// ShortToken generates a less secure token that can be easily shared.
func ShortToken() (string, error) {
	return generate(4)
}

// Token generates a secure token for authentication.
func Token() (string, error) {
	return generate(32)
}
