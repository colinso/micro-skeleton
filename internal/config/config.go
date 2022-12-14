package config

const (
	envDev = "dev"
)

type Config struct {
	AppEnv  string
	AppName string
	Host    string
	Port    string
}

// SetEnvs is where you should bind env keys to a config field
func (c Config) SetEnvs(setter *configSetter) {
	setter.SetStringEnv("APP_NAME", &c.AppName)
	setter.SetStringEnv("HOST", &c.Host)
	setter.SetStringEnv("PORT", &c.Port)
}

// NewConfig will set up config initialization and return a new, populated config
func NewConfig() *Config {
	cfg := Config{}
	setter := configSetter{}

	// Load APP_ENV first, this always must be explicitly set
	setter.SetStringEnv("APP_ENV", &cfg.AppEnv)
	if cfg.AppEnv == envDev {
		cfg = defaultConfig
	}

	// Set Configs from env vars
	cfg.SetEnvs(&setter)

	// Panic and print missing envs
	if errs := setter.getErrors(); errs != nil {
		panic(errs.Error())
	}
	return &cfg
}
