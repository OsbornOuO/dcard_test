package configuration

import (
	"dcard/internal/http"
	"dcard/internal/redis"
	"encoding/json"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gitlab.com/howmay/gopher/zlog"
	"go.uber.org/fx"
)

// Configuration 相關服務的設定值
type Configuration struct {
	fx.Out

	App   *App          `mapstructure:"app"`
	Log   *zlog.Config  `mapstructure:"log"`
	Redis *redis.Config `mapstructure:"redis"`
	HTTP  *http.Config  `mapstructure:"http"`
}

// NewInjection 依賴注入
func (c *Configuration) NewInjection() *Configuration {
	return c
}

// New 讀取App 啟動程式設定檔
func New() (*Configuration, error) {
	viper.AutomaticEnv()

	var config Configuration
	configStr := viper.GetString("CONFIG_APP")
	if configStr != "" {
		if err := json.Unmarshal([]byte(configStr), &config); err != nil {
			log.Err(err).Msgf("Fail to unmarshal config env to struct, env: %s\n", configStr)
		} else {
			if viper.GetString("PORT") != "" {
				config.HTTP.Address = ":" + viper.GetString("PORT")
			}

			if viper.GetString("REDISCLOUD_URL") != "" {
				config.Redis.Address = ":" + viper.GetString("REDISCLOUD_URL")
			}

			return &config, nil
		}
	}

	configPath := viper.GetString("CONFIG_PATH")
	if configPath == "" {
		configPath = "./deploy/config"
	}

	configName := viper.GetString("CONFIG_NAME")
	if configName == "" {
		configName = "app"
	}

	if projDIR := viper.GetString("PROJ_DIR"); projDIR != "" {
		configPath = strings.ReplaceAll(configPath, ".", projDIR)
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Msgf("Error reading config file, %s", err)
		return &config, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Error().Msgf("unable to decode into struct, %v", err)
		return &config, err
	}

	if viper.GetString("PORT") != "" {
		config.HTTP.Address = ":" + viper.GetString("PORT")
	}

	return &config, nil
}
