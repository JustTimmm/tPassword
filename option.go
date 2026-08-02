package tPassword

type Option func(*config)

func WithLowercase() Option {
	return func(c *config) {
		c.includeLowercase = true
	}
}

func WithUppercase() Option {
	return func(c *config) {
		c.includeUppercase = true
	}
}

func WithDigits() Option {
	return func(c *config) {
		c.includeDigits = true
	}
}

func WithSymbols() Option {
	return func(c *config) {
		c.includeSymbols = true
	}
}
