package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nmultiple/autotransfer/discord"
	"github.com/nmultiple/autotransfer/filter"
	"github.com/nmultiple/autotransfer/mail"
)

func main() {
	mailServer := os.Getenv("MXB_MAIL_SERVER")
	mailUser := os.Getenv("MXB_MAIL_USER")
	mailPassword := os.Getenv("MXB_MAIL_PASSWORD")
	discordWebHookUrl := os.Getenv("MXB_DISCORD_WEBHOOK")

	if mailServer == "" || mailUser == "" ||
		mailPassword == "" || discordWebHookUrl == "" {
		log.Fatal("Required environment variables not set")
	}

	textFilter := filter.New()

	mails := mail.Fetch(mailServer, mailUser, mailPassword)
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

		if err := discord.SendMessage(text, discordWebHookUrl); err != nil {
			fmt.Println(err)
		}
	}
}
