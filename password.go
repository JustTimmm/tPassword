package tPassword

import (
	"crypto/rand"
	"math/big"
)

// TODO: Add comments

func Generate(length int, options ...Option) (string, error) {
	cfg := defaultConfig(length)
	for _, option := range options {
		option(&cfg)
	}
	return generateFromConfig(cfg)
}

func GenerateWithCustomCharset(length int, charset string) (string, error) {
	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
