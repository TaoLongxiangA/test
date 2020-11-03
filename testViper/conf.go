package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const (
	kAppName       = "APP_NAME"
	kConfigServer  = "CONFIG_SERVER"
	kConfigLabel   = "CONFIG_LABEL"
	kConfigProfile = "CONFIG_PROFILE"
	kConfigType    = "CONFIG_TYPE"
	kAmqpURI       = "AmqpURI"
)

var (
	Resume ResumeConfig
)

type ResumeConfig struct {
	Name   string
	Age    int
	Gender string
}

func init() {
	viper.AutomaticEnv()
	initDefault()

	if err := loadRemoteConfig(); err != nil {
		log.Fatal("Fail to load config", err)
	}

	if err := sub("resume", &Resume); err != nil {
		log.Fatal("Fail to parse config", err)
	}
}

func initDefault() {
	viper.SetDefault(kAppName, "client-demo")
	viper.SetDefault(kConfigServer, "http://localhost:8888")
	viper.SetDefault(kConfigLabel, "master")
	viper.SetDefault(kConfigProfile, "dev")
	viper.SetDefault(kConfigType, "yaml")
	viper.SetDefault(kAmqpURI, "amqp://admin:admin@114.67.98.210:5672")
}

func loadRemoteConfig() (err error) {
	confAddr := fmt.Sprintf(
		"%v/%v-%v.yml",
		viper.Get(kConfigServer),
		viper.Get(kAppName),
		viper.Get(kConfigProfile),
	)
	resp, err := http.Get(confAddr)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	viper.SetConfigType(viper.GetString(kConfigType))
	//log.Println(viper.GetString(kConfigType))
	if err = viper.ReadConfig(resp.Body); err != nil {
		return
	}
	log.Println("Load config from: ", confAddr)
	return
}

func sub(key string, value interface{}) error {
	log.Printf("配置文件的前缀为：%v", key)
	sub := viper.Sub(key)
	//log.Println("sub above")
	sub.AutomaticEnv()
	//log.Println("sub below")
	sub.SetEnvPrefix(key)
	return sub.Unmarshal(value)
}
