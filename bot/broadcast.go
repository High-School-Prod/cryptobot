package bot

import (
	tb "github.com/tucnak/telebot"
)

func SendEvent(event string) {
	events <- event
}

func broadcastEvents() {
	for event := range events {
		for _, client := range activeClients {
			bot.Send(client, event, tb.ParseMode("HTML"))
		}
	}
}
