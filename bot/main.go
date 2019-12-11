package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	k0 := "AAFGFSqkW6dUqAo2kJ"
	k1 := "821888966"
	l10 := "3bDA06jkpUxzdlcaM"
	key := k1 + ":" + k0 + l10
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	_, err = bot.RemoveWebhook()
	if err != nil {
		log.Println("Cant remove webhook")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		if update.CallbackQuery == nil {
			msg := tgbotapi.NewMessage(conv, "Yo! Life!")
			fmt.Println("kek3")
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonURL("live game", "https://crystal-factory.web.app/#/"),
				),
			)
			fmt.Println("kek4")

			bot.Send(msg)
			continue
		}

		fmt.Println("kek")
		// chID := update.CallbackQuery.ChatInstance
		// chID := update.CallbackQuery.From.ID
		bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert: false, Text: "Play", URL: "https://crystal-factory.firebaseapp.com/"})
		// https://crystal-factory.web.app/#/

		// chID := update.
		// fmt.Println(chID)
		// fmt.Println("kek2")
		// conv, err := strconv.ParseInt(strconv.Itoa(chID), 10, 64)
		// if err != nil {
		// 	panic(err)
		// }
		// msg := tgbotapi.NewMessage(conv, "Yo! Life!")
		// fmt.Println("kek3")
		// // msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		// // 	tgbotapi.NewInlineKeyboardRow(
		// // 		tgbotapi.InlineKeyboardButton{CallbackGame: &tgbotapi.CallbackGame{}},
		// // 	),
		// // )
		// msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		// 	tgbotapi.NewInlineKeyboardRow(
		// 		tgbotapi.NewInlineKeyboardButtonURL("live game", "https://crystal-factory.web.app/#/"),
		// 	),
		// )
		// fmt.Println("kek4")

		// bot.Send(msg)

		// if update.CallbackQuery != nil {
		// 	fmt.Println(update.CallbackQuery)
		// 	continue
		// }

		// bot.GetGameHighScores

		// if update.Message != nil {
		// 	if update.Message.IsCommand() {

		// 		fmt.Println(update.Message)
		// 	} else if update.Message.Text != "" {
		// 		fmt.Println(update.Message)
		// 	}
		// 	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		// 	continue
		// }
	}
}
