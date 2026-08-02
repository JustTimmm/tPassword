package tpassword

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

	return "", nil
}
