package config

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	ConfigName = "config"
	ConfigType = "yaml"
)

var appConfig Config

func Init(ctx context.Context, path, configName, envPrefix string) error {
	viper.AddConfigPath(path)
	// try to look for config in current working directory
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)
	viper.SetConfigType(ConfigType)

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "Config.Init")
	}

	// Auto read config values from env
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// Initialize global config variable from yaml config file
	if err := viper.Unmarshal(&appConfig); err != nil {
		return errors.Wrap(err, "config.Init")
	}
	return nil
}

func Get() Config {
	return appConfig
}
