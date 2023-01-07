package config

import (
	"log"

	"github.com/alexflint/go-arg"
)

var Team = make(map[int64]string)

var Args struct {
	LOG_LEVEL        string `arg:"required,env"`
	TG_BOT_KEY       string `arg:"required,env"`
	TG_ADMIN_CHAT_ID int64  `arg:"required,env"`
	TG_API_URL       string `arg:"required,env"`
	TG_NIKITA_ID     int64  `arg:"required,env"`
	TG_ARTEM_ID      int64  `arg:"required,env"`
	TG_MAKS_ID       int64  `arg:"required,env"`
	TG_ANTON_ID      int64  `arg:"required,env"`
	TG_EGOR_ID       int64  `arg:"required,env"`
	TG_KATE_ID       int64  `arg:"required,env"`
	Team             map[int64]string
}

func Validate() {
	if err := arg.Parse(&Args); err != nil {
		log.Fatal(err)
	}

	Team[Args.TG_NIKITA_ID] = "Никита"
	Team[Args.TG_ARTEM_ID] = "Артем"
	Team[Args.TG_ANTON_ID] = "Антон"
	Team[Args.TG_MAKS_ID] = "Макс"
	Team[Args.TG_KATE_ID] = "Катя"
	Team[Args.TG_EGOR_ID] = "Егор"
}
