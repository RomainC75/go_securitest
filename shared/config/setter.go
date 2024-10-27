package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.AutomaticEnv()

	for _, v := range dbVars {
		if !viper.IsSet(v) {
			log.Fatalf("Environment variable %s not set", v)
		}
	}

}
