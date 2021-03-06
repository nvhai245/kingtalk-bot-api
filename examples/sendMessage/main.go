package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
)

var testPeerID int32 = 1736098

func main() {
	bot, err := ktbotapi.NewBotAPI("332000:YHE508a4V8HL8yKjlHwchA6RXYEjLpIIXExUzwaH")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.FirstName)

	var entities []ktbotapi.MessageEntity
	entities = append(entities, ktbotapi.MessageEntity{
		Type:   "mention",
		Offset: 0,
		Length: 10,
		URL:    "",
		User:   nil,
	})

	msg := ktbotapi.MessageConfig{
		BaseChat: ktbotapi.BaseChat{
			PeerID:              testPeerID,
			ChannelUsername:     "",
			ReplyToMessageID:    0,
			ReplyMarkup:         nil,
			DisableNotification: false,
		},
		Text:                  "@hainv_bot",
		ParseMode:             "",
		DisableWebPagePreview: false,
		Entities:              entities,
	}
	bot.Send(msg)
}
