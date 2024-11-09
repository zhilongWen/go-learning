package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Admin Admin
}

type Admin struct {
	Username string
	Password string
	Path     string
	DbName   string
	Config   string
}

var Dbconfig Config

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("/Users/wenzhilong/warehouse/space/go-learning/gin-admin/test-server/config/dbconfig")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := v.ReadInConfig(); err != nil {
			fmt.Println("Error reading config file:", err)
		}
	})
	if err := v.Unmarshal(&Dbconfig); err != nil {
		fmt.Println("Error unmarshalling config file:", err)
	}
}
