package config

import "github.com/spf13/viper"

type EmailSMTP struct {
	Email           string `mapstructure:"EMAIL_HOST_USER"`
	Password        string `mapstructure:"EMAIL_HOST_PASSWORD"`
	EmailHost       string `mapstructure:"EMAIL_HOST"`
	EmailPortString string `mapstructure:"EMAIL_PORT"`
	EmailPort       int    `mapstructure:"EMAIL_PORT"`
}

func LoadConfig(path string) (config *EmailSMTP, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return config, nil
}
