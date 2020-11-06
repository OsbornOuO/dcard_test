package configuration

// App ..
type App struct {
	IPRatelimitingCount int `mapstructure:"ip_rate_limiting_count"`
	IPRatelimitingSec   int `mapstructure:"ip_rate_limiting_sec"`
}
