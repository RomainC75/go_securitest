package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("../shared/config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("err : ", err.Error())
		log.Fatal("error reading config")
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}

}
