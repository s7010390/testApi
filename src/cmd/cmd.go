package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/s7010390/testApi/config"
	"github.com/s7010390/testApi/interface/http"
	"github.com/s7010390/testApi/logger"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

var wg sync.WaitGroup
var configFile string

func initConfig() {
	// Init viper
	// Init viper
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./config")
		viper.SetConfigType("yaml")
	}
	// Read Config
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "unable to read config: %v\n", err)
		os.Exit(1)
	}
	// Parse Json Decimal to Float
	decimal.MarshalJSONWithoutQuotes = true

	config.LoadConfig()
	// config timezone
	timeZone, err := time.LoadLocation(viper.GetString("System.TimeZone"))
	if err != nil {
		panic("Not found timezone: " + viper.GetString("System.TimeZone") + " please check config SystemConfig.TimeZone\n" + err.Error())
	}
	time.Local = timeZone
}

func initListenInterface() {
	wg.Add(1)
	go func() {
		http.InitHttpServer()
		wg.Done()
	}()
}

func initListenOsSignal() {
	wg.Add(1)
	go func() {
		var count int
		chanOsSignal := make(chan os.Signal, 2)
		signal.Notify(chanOsSignal, syscall.SIGTERM, os.Interrupt)

		go func() {
			for getSignal := range chanOsSignal {
				if getSignal == os.Interrupt || getSignal == syscall.SIGTERM {
					count++
					if count == 2 {
						logger.Logger.Info("Forcefully exiting")
						os.Exit(1)
					}

					go func() {
						http.ShutdownHttpServer()
					}()

					logger.Logger.Info("Signal SIGKILL caught. shutting down")
					logger.Logger.Info("Catching SIGKILL one more time will forcefully exit")

					wg.Done()
				}
			}
			close(chanOsSignal)
		}()
	}()
}
