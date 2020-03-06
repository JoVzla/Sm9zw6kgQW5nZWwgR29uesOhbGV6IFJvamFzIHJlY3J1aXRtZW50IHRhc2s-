package main

import (
	"fmt"
	"gogoapps/logger"
	"gogoapps/models/config"
	"gogoapps/server"
	"os"
	"os/signal"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
)

var cach *cache.Cache
var configuration config.Configuration

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Log.Fatalf("unable to decode into struct, %v", err)
	}
	logger.Log.Infof("port for this application is %s", configuration.Port)
	logger.Log.Infof("api key is %s", configuration.ApiKey)

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed, make logic")
	})

	viper.WatchConfig()
}

func main() {

	initCache()

	//Start rest server
	go server.Start(cach, configuration)

	logger.Log.Info("Servers Started")

	// Wait for Ctrl + C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch
	logger.Log.Info("Stopping servers...")
	server.Stop()
}

func initCache() {
	// purges expired items every 1 minutes
	cach = cache.New(30*time.Second, 1*time.Minute)
}
