package tpassword

// TODO: Add comments

func Generate(length int, options ...Option) (string, error) {
	cfg := defaultConfig(length)
	for _, option := range options {
		option(&cfg)
	}
	return generateFromConfig(cfg)
}
