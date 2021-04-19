package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var acceptedUsers = []int{455753154, 419530579}

func isUserValid(update *tb.Update) bool {
	if update.Message == nil {
		return true
	}

	for _, id := range acceptedUsers {
		if id == update.Message.Sender.ID {
			return true
		}
	}
	bot.Send(update.Message.Sender, "Sorry, you are not authorized to use this bot.")
	return false
}
