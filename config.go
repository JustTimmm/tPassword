package tPassword

type config struct {
	length           int
	includeLowercase bool
	includeUppercase bool
	includeDigits    bool
	includeSymbols   bool
}

func defaultConfig(length int) config {
	return config{
		length:           length,
		includeLowercase: true,
		includeDigits:    true,
		includeUppercase: false,
		includeSymbols:   false,
	}
}
