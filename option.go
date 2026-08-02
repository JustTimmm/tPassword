package tpassword

// TODO: Add comments

type Option func(*config)

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
