package configs

import (
	"flag"
	"github.com/spf13/viper"
	"log/slog"
)

// Configurator dummy
type Configurator struct {
	configPath string
}

// NewConfigurator create new Configurator
func NewConfigurator() *Configurator {
	c := &Configurator{}
	flag.StringVar(&c.configPath, "config", "", "specify config path")
	c.load()
	return c
}

// Init env data from files (default: .env, .env.local)
func (c *Configurator) load() {
	if c.configPath == "" {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("error loading config file", slog.Any("err", err))
		panic("error loading config file")
	}

	err := viper.MergeInConfig()
	if err != nil {
		slog.Error("error merge config file", slog.Any("err", err))
		return
	}
}
