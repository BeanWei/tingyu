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
		Address string `yaml:"address"`
		IsDev   bool   `yaml:"is_dev"`
	} `yaml:"server"`
	JWT struct {
		SecretKey   string `yaml:"secret_key"`
		TimeoutDays int    `yaml:"timeout_days"`
	}
	Database struct {
		Read  string `yaml:"read"`
		Write string `yaml:"write"`
	} `yaml:"database"`
	Redis struct {
		Addr string `toml:"addr"`
		Pass string `toml:"pass"`
		DB   int    `toml:"db"`
	} `toml:"redis"`
}

func Cfg() *conf {
	cfgOnce.Do(func() {
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		cfgfile := viper.GetString("c")
		if cfgfile == "" {
			panic("config file is missing")
		}
		viper.SetConfigFile(cfgfile)

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("read config faild: %+v", err))
		} else {
			fmt.Println("Using config file: ", viper.ConfigFileUsed())
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(fmt.Errorf("viper unmarshal config faild: %+v", err))
		}
	})
	return cfg
}
