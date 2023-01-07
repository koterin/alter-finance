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

func OnStart() tb.HandlerFunc {
	return func(c tb.Context) error {
		if !c.Message().Private() {
			log.Error("Error: chat is not private with user ", c.Sender().Username)

			return c.Send(entity.TextInternalError)
		}

		log.Info("User with name = ", c.Message().Sender.Username, " and ID = ", strconv.Itoa(int(c.Message().Chat.ID)), " started bot")

		name, present := config.Team[c.Message().Chat.ID]
		if !present {
			log.Error("No user in Team")
			return c.Send(entity.TextUnkownUser)
		}
		return c.Send(fmt.Sprintf(entity.TextInstructions, name))
	}
}

func OnText() tb.HandlerFunc {
	return func(c tb.Context) error {
		log.Info("User with name = ", c.Message().Sender.Username, " and ID = ", strconv.Itoa(int(c.Message().Chat.ID)), " sent message: ", c.Message().Text)

		name, present := config.Team[c.Message().Chat.ID]
		if !present {
			log.Info("User is not recognized")
			return c.Send(entity.TextUnkownUser)
		}

		usecase.SendNewRecord(name, c.Message().Text)

		return c.Send(entity.TextReceived)
	}
}
