package tpassword

import "errors"

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

	// list of all characters define by the options of the lib
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

	// management of errors
	if cfg.length <= 0 {
		return "", errors.New("you try to generate a empty password")
	}
	if charset == "" {
		return "", errors.New("no one charset are selected")
	}

	return "", nil
}
