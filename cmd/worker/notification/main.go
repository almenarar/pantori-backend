package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pantori/internal/domains/notifiers"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("/go/bin/config.json")
	viper.ReadInConfig()

	notifier := notifiers.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		log.Println("Shutting down gracefully...")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			notifier.NotifyExpiredGoods()
			time.Sleep(24 * time.Hour)
		}
	}
}
