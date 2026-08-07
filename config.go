package tPassword

// TODO: Add comments

type config struct {
	length    int
	lowercase bool
	uppercase bool
	digits    bool
	symbols   bool
	ambiguous bool
}

func defaultConfig(length int) config {
	return config{
		length:    length,
		lowercase: true,
		digits:    true,
		uppercase: false,
		symbols:   false,
		ambiguous: false,
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
	if cfg.ambiguous {
		charset = removeChars(charset, "O01lI")
	}

	// error handling
	if cfg.length <= 0 {
		return "", ErrInvalidLength
	}
	if charset == "" {
		return "", ErrEmptyCharset
	}

	// pick randoms characters using crypto/rand
	result, err := randomStringFromCharset(cfg.length, charset)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
