package controller

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"

	"alter-finance/config"
	"alter-finance/internal/entity"
	"alter-finance/internal/usecase"
)

var (
	team = map[int64]string{
		config.Args.TG_NIKITA_ID: "Никита",
		config.Args.TG_ARTEM_ID:  "Артем",
		config.Args.TG_ANTON_ID:  "Антон",
		config.Args.TG_MAKS_ID:   "Макс",
		config.Args.TG_KATE_ID:   "Катя",
		config.Args.TG_EGOR_ID:   "Егор",
	}
)

func OnStart() tb.HandlerFunc {
	return func(c tb.Context) error {
		if !c.Message().Private() {
			log.Error("Error: chat is not private with user ", c.Sender().Username)

			return c.Send(entity.TextInternalError)
		}

		name, present := team[c.Message().Chat.ID]
		if !present {
			return c.Send(entity.TextUnkownUser)
		}
		return c.Send(fmt.Sprintf(entity.TextInstructions, name))
	}
}

func OnText() tb.HandlerFunc {
	return func(c tb.Context) error {
		chatId := strconv.Itoa(int(c.Message().Chat.ID))

		log.Info("User ", chatId, " sent message: ", c.Message().Text)
		usecase.SendNewRecord(team[c.Message().Chat.ID], c.Message().Text)

		return c.Send(entity.TextReceived)
	}
}