package bot

import (
	"fmt"
	"github.com/NazarNintendo/cryptobot/fetchers"
	"github.com/NazarNintendo/cryptobot/model"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

var activeClients = make(map[int]*model.Client)
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

	//bot.Handle("/start", func(m *tb.Message) {
	//	bot.Send(m.Sender, "Hello! You've subscribed to Crypto Monitoring")
	//	activeClients[m.Sender.ID] = &model.Client{User: m.Sender}
	//})
	//
	//bot.Handle("/stop", func(m *tb.Message) {
	//	bot.Send(m.Sender, "You've unsubscribed from Crypto Monitoring. Bye!")
	//	delete(activeClients, m.Sender.ID)
	//})

	bot.Handle("/help", func(m *tb.Message) {
		bot.Send(m.Sender, "My name is Crypto Deathrow and I will gather cryptocurrency data for you.\n\n"+
			"/info          ➙   fetch real-time data about cryptocurrencies.\n"+
			"/settings   ➙   configure settings.\n")
	})

	bot.Handle("/info", func(m *tb.Message) {
		bot.Send(m.Sender, fetchers.OneTimeFetch(), tb.ParseMode("HTML"))
	})

	bot.Handle("/settings", func(m *tb.Message) {
		bot.Send(m.Sender, "Hewllo!", mainMenu)
	})

	bot.Handle(&monitoringsBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		bot.Edit(c.Message, "Hewllo!", makeDecisionMenu("monitoring"))
	})

	bot.Handle(&backBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		bot.Edit(c.Message, "Hewllo!", mainMenu)
	})

	bot.Handle(&newBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		bot.Edit(c.Message, "Choose Cryptocurrency", currenciesMenu)
	})

	bot.Handle(&deleteBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{Text: "All monitorings successfully deleted!"})
		deleteMonitorings(c.Sender)
		bot.Edit(c.Message, "Hewllo!", mainMenu)
	})

	bot.Handle(&dogeBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		bot.Edit(c.Message, fmt.Sprintf("How often for %v", c.Data), makeFrequencyMenu(c.Data))
	})

	bot.Handle(&btcBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		bot.Edit(c.Message, fmt.Sprintf("How often for %v", c.Data), makeFrequencyMenu(c.Data))
	})

	bot.Handle(&freq10mBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		message, confirmationMenu := makeConfirmationMenu(c.Data)
		bot.Edit(c.Message, message, confirmationMenu)
	})

	bot.Handle(&freq30mBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		message, confirmationMenu := makeConfirmationMenu(c.Data)
		bot.Edit(c.Message, message, confirmationMenu)
	})

	bot.Handle(&freq1hBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		message, confirmationMenu := makeConfirmationMenu(c.Data)
		bot.Edit(c.Message, message, confirmationMenu)
	})

	bot.Handle(&freq2hBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		message, confirmationMenu := makeConfirmationMenu(c.Data)
		bot.Edit(c.Message, message, confirmationMenu)
	})

	bot.Handle(&freq6hBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		message, confirmationMenu := makeConfirmationMenu(c.Data)
		bot.Edit(c.Message, message, confirmationMenu)
	})

	bot.Handle(&freq12hBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		message, confirmationMenu := makeConfirmationMenu(c.Data)
		bot.Edit(c.Message, message, confirmationMenu)
	})

	bot.Handle(&freq24hBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{})
		message, confirmationMenu := makeConfirmationMenu(c.Data)
		bot.Edit(c.Message, message, confirmationMenu)
	})

	bot.Handle(&okBtn, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{Text: "Monitoring setup successful!"})
		bot.Delete(c.Message)
		addMonitoring(c.Sender, c.Data)
	})

	bot.Handle(&metricsBtn, func(c *tb.Callback) {
		// ...
		// Always respond!
		bot.Respond(c, &tb.CallbackResponse{})
		bot.Send(c.Sender, "Next")
	})

	go broadcastEvents()
	bot.Start()
}
