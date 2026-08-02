package tPassword

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// TODO: Add comments

type config struct {
	length    int
	lowercase bool
	uppercase bool
	digits    bool
	symbols   bool
}

func defaultConfig(length int) config {
	return config{
		length:    length,
		lowercase: true,
		digits:    true,
		uppercase: false,
		symbols:   false,
	}
}

func generateFromConfig(cfg config) (string, error) {
	var charset string

	// list of characters allowed by the config options
	if cfg.lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if cfg.uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if cfg.digits {
		charset += "0123456789"
	}
	if cfg.symbols {
		charset += ".?!&#()$%+-=_-@"
	}

	// error handling
	if cfg.length <= 0 {
		return "", errors.New("cannot generate an empty password")
	}
	if charset == "" {
		return "", errors.New("no charset selected")
	}

	// pick randoms characters using crypto/rand
	result := make([]byte, cfg.length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
