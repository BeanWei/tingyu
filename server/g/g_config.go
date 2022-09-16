package g

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg     *conf
	cfgOnce sync.Once
)

type conf struct {
	Server struct {
		Address string `mapstructure:"address"`
		IsDev   bool   `mapstructure:"is_dev"`
	} `mapstructure:"server"`
	JWT struct {
		SecretKey   string `mapstructure:"secret_key"`
		TimeoutDays int    `mapstructure:"timeout_days"`
	} `mapstructure:"jwt"`
	Database struct {
		Read  string `mapstructure:"read"`
		Write string `mapstructure:"write"`
	} `mapstructure:"database"`
	Redis struct {
		Addr string `mapstructure:"addr"`
		Pass string `mapstructure:"pass"`
		DB   int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	Meilisearch struct {
		Host   string `mapstructure:"host"`
		APIKey string `mapstructure:"api_key"`
	} `mapstructure:"meilisearch"`
	Operation struct {
		Audit bool `mapstructure:"audit"`
	} `mapstructure:"operation"`
}

func Cfg() *conf {
	cfgOnce.Do(func() {
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		cfgFile := viper.GetString("ty_cfg_file")
		if cfgFile == "" {
			panic("config file is missing")
		}
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("read config failed: %+v", err))
		} else {
			fmt.Println("Using config file: ", viper.ConfigFileUsed())
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(fmt.Errorf("viper unmarshal config failed: %+v", err))
		}
	})
	return cfg
}
