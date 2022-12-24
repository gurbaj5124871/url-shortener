package config

import (
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var logger zerolog.Logger = log.With().Str("logger", "config").Logger()

type Server struct {
	Port string `mapstructure:"port"`
}

type Ignite struct {
	Host     string `mapstructure:"host"`
	Post     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	TLS      bool   `mapstructure:"tls"`
}

type Config struct {
	Server Server
	Ignite Ignite
}

var config *Config

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("config/")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("error on parsing configuration file")
	}
	_ = viper.Unmarshal(&config)
}

func RelativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *Config {
	return config
}
