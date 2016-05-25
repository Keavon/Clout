package token

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"strconv"
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

// GameID generates a less secure token that can be easily shared.
func GameID() (string, error) {
	token := ""

	// 5 character long string
	for i := 0; i < 5; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		token += strconv.Itoa(int(num.Int64()))
	}
	return token, nil
}

// Token generates a secure token for authentication.
func Token() (string, error) {
	return generate(32)
}
