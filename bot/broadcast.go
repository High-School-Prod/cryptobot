package bot

import (
	"fmt"
	"github.com/NazarNintendo/cryptobot/model"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func SendEvent(event string) {
	events <- event
}

func broadcastEvents() {
	for event := range events {
		for _, client := range activeClients {
			bot.Send(client.User, event, tb.ParseMode("HTML"))
		}
	}
}

func NotifyClients(aggregationCounter int, cryptos []model.Crypto) {
	for _, client := range activeClients {
		for _, monitoring := range client.Monitorings {
			switch monitoring.Frequency {
			case 1:
				if aggregationCounter%1 == 0 {
					sendNotification(client, monitoring.Currency, cryptos)
				}
			case 2:
				if aggregationCounter%3 == 0 {
					sendNotification(client, monitoring.Currency, cryptos)
				}
			case 3:
				if aggregationCounter%6 == 0 {
					sendNotification(client, monitoring.Currency, cryptos)
				}
			case 4:
				if aggregationCounter%12 == 0 {
					sendNotification(client, monitoring.Currency, cryptos)
				}
			case 5:
				if aggregationCounter%36 == 0 {
					sendNotification(client, monitoring.Currency, cryptos)
				}
			case 6:
				if aggregationCounter%72 == 0 {
					sendNotification(client, monitoring.Currency, cryptos)
				}
			case 7:
				if aggregationCounter%144 == 0 {
					sendNotification(client, monitoring.Currency, cryptos)
				}
			}
		}
	}

}

func sendNotification(client *model.Client, currencyId int, cryptos []model.Crypto) {
	currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
	timestamp := fmt.Sprintf("<i>%v</i>\n", currentTime)
	switch currencyId {
	case 1:
		bot.Send(client.User, cryptos[0].String()+timestamp, tb.ParseMode("HTML"))
	case 74:
		bot.Send(client.User, cryptos[1].String()+timestamp, tb.ParseMode("HTML"))
	}
}
