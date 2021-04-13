package bot

import (
	"../fetchers"
	tb "github.com/tucnak/telebot"
	"log"
	"os"
	"time"
)

var activeClients = make(map[int]*tb.User)
var events = make(chan string)

var bot *tb.Bot

func CreateBot() {
	poller := &tb.LongPoller{Timeout: 10 * time.Second}
	authProtected := tb.NewMiddlewarePoller(poller, isUserValid)

	var err error
	bot, err = tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: authProtected,
	})

	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Hello! You've subscribed to Crypto Monitoring")
		activeClients[m.Sender.ID] = m.Sender
	})

	bot.Handle("/stop", func(m *tb.Message) {
		bot.Send(m.Sender, "You've unsubscribed from Crypto Monitoring. Bye!")
		delete(activeClients, m.Sender.ID)
	})

	bot.Handle("/help", func(m *tb.Message) {
		bot.Send(m.Sender, "My name is Crypto Deathrow and I will gather cryptocurrency data for you.\n\n"+
			"/start        ➙   subscribe to events. They are sent as messages every twelve hours.\n"+
			"/stop         ➙   unsubscribe from event messaging.\n"+
			"/info          ➙   manually get real-time data about cryptocurrencies.\n"+
			"/metrics   ➙   configure metric alarm monitoring.\n")
	})

	bot.Handle("/info", func(m *tb.Message) {
		bot.Send(m.Sender, fetchers.OneTimeFetch(), tb.ParseMode("HTML"))
	})

	go broadcastEvents()
	bot.Start()
}
