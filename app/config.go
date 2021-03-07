package app

import "github.com/spf13/viper"

type Config struct {
	*viper.Viper
}

func InitConfig(defaultConfigName string) (*Config, error) {
	config := &Config{viper.New()}
	viper.SetEnvPrefix("calculator")
	var env string
	if err := viper.BindEnv("env"); err != nil {
		env = defaultConfigName
	} else {
		env = viper.GetString("env")
		if env == "" {
			env = defaultConfigName
		}
	}
	config.SetConfigName(env)
	config.AddConfigPath("./config")
	config.AddConfigPath("$HOME/.config/calculator")
	config.SetEnvPrefix("calculator")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		if env != "" && env != defaultConfigName {
			config.SetConfigName(defaultConfigName)
			if err := config.ReadInConfig(); err != nil {
				return nil, err
			}
			config.WatchConfig()
			return config, nil
		}
		return nil, err
	}
	config.WatchConfig()
	return config, nil
}
