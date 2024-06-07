package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	notifier.NotifyExpiredGoods(ctx)

	<-ctx.Done()
	log.Println("Job runner stopped.")
}
