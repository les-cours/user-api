package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort        string
	UserService     *UserServiceConfig
	SettingsService *SettingsServiceConfig
	VisitorService  *VisitorServiceConfig
}

type UserServiceConfig struct {
	Host string
	Port string
}

type SettingsServiceConfig struct {
	Host string
	Port string
}

type VisitorServiceConfig struct {
	Host string
	Port string
}

var Conf *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("USERS_SERVICE_HOST")
	viper.BindEnv("USERS_SERVICE_PORT")

	viper.BindEnv("SETTINGS_SERVICE_HOST")
	viper.BindEnv("SETTINGS_SERVICE_PORT")

	Conf = &Config{
		HttpPort: viper.GetString("HTTP_PORT"),
		UserService: &UserServiceConfig{
			Host: viper.GetString("USERS_SERVICE_HOST"),
			Port: viper.GetString("USERS_SERVICE_PORT"),
		},
		SettingsService: &SettingsServiceConfig{
			Host: viper.GetString("SETTINGS_SERVICE_HOST"),
			Port: viper.GetString("SETTINGS_SERVICE_PORT"),
		},
		VisitorService: &VisitorServiceConfig{
			Host: viper.GetString("VISITOR_SERVICE_HOST"),
			Port: viper.GetString("VISITOR_SERVICE_PORT"),
		},
	}
}
