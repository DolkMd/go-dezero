package conf

type (
	config struct {
		EnableBackprop bool
	}
	ConfigOption func(*config)
)

var Config = config{
	EnableBackprop: true,
}

func UsingConfig(fn func(), opts ...ConfigOption) {
	copiedConf := Config
	defer func() { Config = copiedConf }()

	for _, opt := range opts {
		opt(&Config)
	}

	fn()
}

func EnableBackprop(enable bool) ConfigOption {
	return func(c *config) {
		c.EnableBackprop = enable
	}
}
