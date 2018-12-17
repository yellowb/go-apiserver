package config

import (
	"github.com/lexkong/log"
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

	// Init config
	if err := c.initConfig(); err != nil {
		return err
	}

	// Init logger
	if err := c.initLog(); err != nil {
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

func (c *Config) initLog() error {
	passLagerCfg := log.PassLagerCfg{
		Writers:viper.GetString("log.writers"),
		LoggerLevel:viper.GetString("log.logger_level"),
		LoggerFile:viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	return log.InitWithConfig(&passLagerCfg)
}


