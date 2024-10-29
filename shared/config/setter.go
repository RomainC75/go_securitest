package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set(varList []ConfigVar) {
	viper.AutomaticEnv()

	for _, v := range varList {
		if !viper.IsSet(string(v)) {
			log.Fatalf("Environment variable %s not set", v)
		}
	}
}
