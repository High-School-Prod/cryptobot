package bot

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

var (
	mainMenu       = &tb.ReplyMarkup{}
	monitoringsBtn = mainMenu.Data("Monitorings", "monitorings")
	metricsBtn     = mainMenu.Data("Metrics", "metrics")
)

var (
	decisionMenu = &tb.ReplyMarkup{}
	newBtn       = decisionMenu.Data("New", "new")
	deleteBtn    = decisionMenu.Data("Delete", "delete")
)

var (
	currenciesMenu = &tb.ReplyMarkup{}
	dogeBtn        = currenciesMenu.Data("DOGE", "doge", "DOGE")
	btcBtn         = currenciesMenu.Data("BTC", "btc", "BTC")
	backBtn        = currenciesMenu.Data("Back", "back", "back")
)

var (
	frequenciesMenu = &tb.ReplyMarkup{}
	freq10mBtn      = frequenciesMenu.Data("10m", "10m", "10m|")
	freq30mBtn      = frequenciesMenu.Data("30m", "30m", "30m|")
	freq1hBtn       = frequenciesMenu.Data("1h", "1h", "1h|")
	freq2hBtn       = frequenciesMenu.Data("2h", "2h", "2h|")
	freq6hBtn       = frequenciesMenu.Data("6h", "6h", "6h|")
	freq12hBtn      = frequenciesMenu.Data("12h", "12h", "12h|")
	freq24hBtn      = frequenciesMenu.Data("24h", "24h", "24h|")
)

var (
	confirmationMenu = &tb.ReplyMarkup{}
	okBtn            = confirmationMenu.Data("Yes", "yes")
	cancelBtn        = confirmationMenu.Data("Cancel", "cancel")
)

func init() {
	mainMenu.Inline(mainMenu.Row(monitoringsBtn, metricsBtn))
	currenciesMenu.Inline(currenciesMenu.Row(dogeBtn, btcBtn), currenciesMenu.Row(backBtn))
}

func makeFrequencyMenu(cryptoName string) *tb.ReplyMarkup {
	frequenciesMenu = &tb.ReplyMarkup{}
	freq10mBtn = frequenciesMenu.Data("10m", "10m", "10m|"+cryptoName)
	freq30mBtn = frequenciesMenu.Data("30m", "30m", "30m|"+cryptoName)
	freq1hBtn = frequenciesMenu.Data("1h", "1h", "1h|"+cryptoName)
	freq2hBtn = frequenciesMenu.Data("2h", "2h", "2h|"+cryptoName)
	freq6hBtn = frequenciesMenu.Data("6h", "6h", "6h|"+cryptoName)
	freq12hBtn = frequenciesMenu.Data("12h", "12h", "12h|"+cryptoName)
	freq24hBtn = frequenciesMenu.Data("24h", "24h", "24h|"+cryptoName)
	frequenciesMenu.Inline(
		frequenciesMenu.Row(
			freq10mBtn,
			freq30mBtn,
			freq1hBtn,
			freq2hBtn,
			freq6hBtn,
			freq12hBtn,
			freq24hBtn,
		),
		frequenciesMenu.Row(backBtn),
	)
	return frequenciesMenu
}

func makeConfirmationMenu(monitoringConfig string) (string, *tb.ReplyMarkup) {
	crypto := strings.Split(monitoringConfig, "|")[1]
	frequency := strings.Split(monitoringConfig, "|")[0]
	message := fmt.Sprintf("Setting up monitoring for %v every %v?", crypto, frequency)
	confirmationMenu = &tb.ReplyMarkup{}
	okBtn = confirmationMenu.Data("Yes", "yes", monitoringConfig)
	confirmationMenu.Inline(confirmationMenu.Row(okBtn, backBtn))
	return message, confirmationMenu
}

func makeDecisionMenu(config string) *tb.ReplyMarkup {
	decisionMenu = &tb.ReplyMarkup{}
	newBtn = decisionMenu.Data("New", "new", config)
	deleteBtn = decisionMenu.Data("Delete", "delete", config)
	decisionMenu.Inline(decisionMenu.Row(newBtn, deleteBtn))
	return decisionMenu
}
