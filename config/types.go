package config

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type Config struct {
	HTTP     *HTTPConfig `mapstructure:"http,omitempty"`
	Database Database    `mapstructure:"database"`
}
type HTTPConfig struct {
	// Hostname
	Host string `mapstructure:"host"`
	// Port
	Port int `mapstructure:"port"`
}
