package bot

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"

	"alter-finance/config"
	"alter-finance/internal/controller"
)

var Bot = &tb.Bot{}

func StartTelegramBot(ctx context.Context) {
	config.Validate()

	level, err := log.ParseLevel(config.Args.LOG_LEVEL)
	if err != nil {
		log.Fatal(err)
	}

	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(level)

	settings := tb.Settings{
		Token: config.Args.TG_BOT_KEY,
		Poller: &tb.LongPoller{
			Timeout: 1 * time.Second,
		},
	}

	Bot, _ = tb.NewBot(settings)

	Bot.Handle("/start", controller.OnStart())
	Bot.Handle(tb.OnText, controller.OnText())

	go func() {
		Bot.Start()
	}()

	log.Info("Alter Finance Bot started")

	<-ctx.Done()

	log.Info("Alter Finance Bot stopped")
	Bot.Stop()
}
