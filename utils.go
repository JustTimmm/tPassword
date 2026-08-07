package tPassword

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func randomStringFromCharset(length int, charset string) (string, error) {
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

func removeChars(charset, toRemove string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(toRemove, r) {
			return -1
		}
		return r
	}, charset)
}
