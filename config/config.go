package config

import (
	"log"

	"github.com/alexflint/go-arg"
)

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
}

func Validate() {
	if err := arg.Parse(&Args); err != nil {
		log.Fatal(err)
	}
}
