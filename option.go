package tPassword

// TODO: Add comments

type Option func(*config)

// Deprecated: Already true in the default config
func WithLowercase() Option {
	return func(c *config) {
		c.lowercase = true
	}
}

func WithUppercase() Option {
	return func(c *config) {
		c.uppercase = true
	}
}

// Deprecated: Already true in the default config
func WithDigits() Option {
	return func(c *config) {
		c.digits = true
	}
}

func WithSymbols() Option {
	return func(c *config) {
		c.symbols = true
	}
}
