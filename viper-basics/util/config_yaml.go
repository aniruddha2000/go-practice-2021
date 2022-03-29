package util

import (
	"github.com/spf13/viper"
)

type ConfigYaml struct {
	Env            string `mapstructure:"env"`
	Consumerbroker string `mapstructure:"consumerbroker"`
	Producerbroker string `mapstructure:"producerbroker"`
	Linetoken      string `mapstructure:"linetoken"`
}

func LoadYamlConfig(path string) (cfg ConfigYaml, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// viper.SetDefault("app.linetoken", "DefaultLineTokenValue")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return
	}
	return
}
