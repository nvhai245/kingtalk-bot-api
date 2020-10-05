package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1735940:93M5fKuPedrQ45PiosUTvl0Z0h5lTrCaw6tYRA9a")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	//updates := bot.ListenForWebhook("/" + bot.Token)
	//go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
	//
	//for update := range updates {
	//	log.Printf("%+v\n", update)
	//}
}