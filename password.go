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
	if length <= 0 {
		return "", ErrInvalidLength
	}
	if charset == "" {
		return "", ErrEmptyCharset
	}

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

func GenerateMulti(count, length int, options ...Option) ([]string, error) {
	passwords := make([]string, count)
	for i := range count {
		password, err := Generate(length, options...)
		if err != nil {
			return nil, err
		}
		passwords[i] = password
	}

	return passwords, nil
}

func GenerateMultiWithCustomCharset(count, length int, charset string) ([]string, error) {
	passwords := make([]string, count)
	for i := range count {
		password, err := GenerateWithCustomCharset(length, charset)
		if err != nil {
			return nil, err
		}
		passwords[i] = password
	}

	return passwords, nil
}
