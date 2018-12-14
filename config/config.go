package config

import (
	"github.com/spf13/viper"
	"strings"
)

const (
	configDir = "conf"
	configFileName = "config"
	configFileType = "yaml"
	envPrefix = "APISERVER"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{Name:cfg}

	//Init config
	if err := c.initConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath(configDir)
		viper.SetConfigName(configFileName)
	}
	viper.SetConfigType(configFileType)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	replacer := strings.NewReplacer(".","_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}



