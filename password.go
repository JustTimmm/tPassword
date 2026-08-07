package tPassword

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

	result, err := randomStringFromCharset(length, charset)
	if err != nil {
		return "", err
	}

	return result, nil
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
