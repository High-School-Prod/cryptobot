package logic

import (
	"github.com/NazarNintendo/cryptobot/bot"
	"github.com/NazarNintendo/cryptobot/fetchers"
	"github.com/NazarNintendo/cryptobot/model"
	"sync"
	"time"
)

var rates = make(chan model.ExchangeRate)
var quotes = make(chan model.CryptoQuote)

func Run() {

	//historicalQuotes := make([]model.CryptoQuote, 72)

	go bot.CreateBot()

	aggregationCounter := 1

	for {
		var waitGroup sync.WaitGroup
		waitGroup.Add(3)

		go fetchers.FetchExchangeRate(rates, &waitGroup)
		go fetchers.FetchLatestQuote(quotes, &waitGroup)

		go func() {
			exchangeRateUSD := <-rates
			quote := <-quotes

			var cryptos []model.Crypto

			cryptos = append(cryptos, formatCrypto(exchangeRateUSD, quote.Data.BTC))
			cryptos = append(cryptos, formatCrypto(exchangeRateUSD, quote.Data.DOGE))

			bot.NotifyClients(aggregationCounter, cryptos)

			//historicalQuotes[aggregationCounter] = quote

			waitGroup.Done()
		}()

		waitGroup.Wait()

		if aggregationCounter == 144 {
			aggregationCounter = 1
		} else {
			aggregationCounter++
		}

		time.Sleep(10 * time.Minute)
	}
}
