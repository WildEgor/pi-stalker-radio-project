package configs

import (
	"github.com/spf13/viper"
	"log/slog"
)

// AppConfig holds the main app configurations
type AppConfig struct {
	Name       string `mapstructure:"name"`
	Mode       string `mapstructure:"mode"`
	HTTPPort   string `mapstructure:"http_port"`
	RPCPort    string `mapstructure:"rpc_port"`
	AssetsPath string `mapstructure:"assets_path"`
}

// NewAppConfig creates app config
func NewAppConfig() *AppConfig {
	cfg := &AppConfig{}

	if err := viper.UnmarshalKey("app", &cfg); err != nil {
		slog.Error("app config parse error")
		panic("logger parse error")
	}

	slog.Info("config", slog.Any("value", cfg))

	return cfg
}

// IsProduction Check is application running in production mode
func (ac *AppConfig) IsProduction() bool {
	return ac.Mode == "production"
}

// IsDebug Check is application running in debug mode
func (ac *AppConfig) IsDebug() bool {
	return ac.Mode == "debug"
}
