package config

import (
	"eicesoft/proxy-api/pkg/env"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var config = new(Config)

type Config struct {
	Server struct {
		Name string `toml:"name"`
		Port string `toml:"port"`
		Cors bool   `toml:"cors"` //是否开启Cors
	} `toml:"Server"`

	MySQL struct {
		Read struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"read"`
		Write struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"write"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"base"`
	} `toml:"mysql"`

	JWT struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"jwt"`
}

func init() {
	viper.SetConfigName(env.Get().Value() + "_config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
}

func Get() Config {
	return *config
}

func ProjectLogFile() string {
	return fmt.Sprintf("./logs/access-%s.log", env.Get().Value())
}
