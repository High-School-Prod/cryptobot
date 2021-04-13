package logic

import (
	"../bot"
	"../fetchers"
	"../model"
	"fmt"
	"sync"
	"time"
)

var rates = make(chan model.ExchangeRate)
var quotes = make(chan model.CryptoQuote)

func Run() {

	historicalQuotes := make([]model.CryptoQuote, 72)

	go bot.CreateBot()

	aggregationCounter := 0

	for {
		var waitGroup sync.WaitGroup
		waitGroup.Add(3)

		go fetchers.FetchExchangeRate(rates, &waitGroup)
		go fetchers.FetchLatestQuote(quotes, &waitGroup)

		go func() {
			exchangeRateUSD := <-rates
			quote := <-quotes

			var cryptos []model.Crypto

			cryptos = append(cryptos, formatCrypto(exchangeRateUSD, quote.Data.BTC, historicalQuotes, aggregationCounter))
			cryptos = append(cryptos, formatCrypto(exchangeRateUSD, quote.Data.DOGE, historicalQuotes, aggregationCounter))

			var message string

			for _, crypto := range cryptos {
				message += crypto.String()
			}

			currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
			message += fmt.Sprintf("<i>%v</i>\n", currentTime)

			if aggregationCounter == 71 {
				bot.SendEvent(message)
			}

			historicalQuotes[aggregationCounter] = quote

			waitGroup.Done()
		}()

		waitGroup.Wait()

		if aggregationCounter == 71 {
			aggregationCounter = 0
		} else {
			aggregationCounter++
		}

		time.Sleep(10 * time.Minute)
	}
}
