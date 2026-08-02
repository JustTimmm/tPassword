package tpassword

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
