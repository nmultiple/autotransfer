package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nmultiple/autotransfer/config"
	"github.com/nmultiple/autotransfer/discord"
	"github.com/nmultiple/autotransfer/filter"
	"github.com/nmultiple/autotransfer/mail"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MailServer         string `envconfig:"mxb_mail_server" required:"true"`
	MailUser           string `envconfig:"mxb_mail_user" required:"true"`
	MailPassword       string `envconfig:"mxb_mail_password" required:"true"`
	DiscordWebHookUrl  string `envconfig:"mxb_discord_webhook" required:"true"`
	FetchIntervalHours int    `envconfig:"mxb_interval_hour" required:"false" default:"3"`
	OneShot            bool   `ignored:"true"`
}

func fetchMailAndPost(cfg Config) {
	textFilter := filter.New()

	mails, err := mail.Fetch(cfg.MailServer, cfg.MailUser, cfg.MailPassword)
	if err != nil {
		log.Println(err)
	}

	for _, m := range mails {
		var messageBuilder strings.Builder

		if len(m.Subject) != 0 {
			messageBuilder.WriteString(m.Subject)
			messageBuilder.WriteString("\n\n")
		}

		text := textFilter.FilterText(m.Text)

		messageBuilder.WriteString(text)

		if m.AttachmentCount != 0 {
			messageBuilder.WriteString("\n\n")
			messageBuilder.WriteString(fmt.Sprintf("%d Attachments.", m.AttachmentCount))
		}

		if err := discord.SendMessage(text, cfg.DiscordWebHookUrl); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		if err := config.LoadDotEnv(); err != nil {
			log.Fatal("Failed to load config:", err)
		}

		if err := envconfig.Process("", &cfg); err != nil {
			log.Fatal("Failed to load config (tried .env):", err)
		}
	}

	if len(os.Args) > 1 && os.Args[1] == "oneshot" {
		cfg.OneShot = true
	}

	if cfg.OneShot {
		fetchMailAndPost(cfg)
	} else {
		for {
			fetchMailAndPost(cfg)
			time.Sleep(time.Hour * time.Duration(cfg.FetchIntervalHours))
		}
	}
}
