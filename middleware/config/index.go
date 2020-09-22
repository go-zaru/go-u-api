package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// ReadConfig to invoke reading of local config
func ReadConfig(cf string) *viper.Viper {
	var instance = viper.New()
	instance.AutomaticEnv()

	if cf != "" {
		instance.SetConfigFile(cf)
		err := instance.ReadInConfig()
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
	// err := instance.Unmarshal(&cfg)
	return instance
}

// ReadRemoteConfig read remote config and return Viper instance
func ReadRemoteConfig(provider string, url string, key string, ext string, cfg interface{}) *viper.Viper {
	var instance = viper.New()
	instance.AutomaticEnv()
	instance.AddRemoteProvider(provider, url, key)
	instance.SetConfigType(ext)
	// read from remote config the first time.
	err := instance.ReadRemoteConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	// unmarshal config
	// instance.Unmarshal(&cfg)
	return instance
}
