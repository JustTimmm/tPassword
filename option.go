package tPassword

// TODO: Add comments

type Option func(*config)

func WithoutLowercase() Option {
	return func(c *config) {
		c.lowercase = false
	}
}

func WithoutDigits() Option {
	return func(c *config) {
		c.digits = false
	}
}

func WithUppercase() Option {
	return func(c *config) {
		c.uppercase = true
	}
}

func WithSymbols() Option {
	return func(c *config) {
		c.symbols = true
	}
}
