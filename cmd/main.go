package main

import (
	"context"
	"os"
	"os/signal"

	"alter-finance/internal/bot"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go bot.StartTelegramBot(ctx)

	<-c
	cancel()
}
