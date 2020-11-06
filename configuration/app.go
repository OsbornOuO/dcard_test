package configuration

// App ..
type App struct {
	IPRatelimitingCount int `json:"ip_rate_limiting_count" mapstructure:"ip_rate_limiting_count"`
	IPRatelimitingSec   int `json:"ip_rate_limiting_sec" mapstructure:"ip_rate_limiting_sec"`
}
