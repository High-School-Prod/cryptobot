package fetchers

import (
	"../model"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func FetchExchangeRate(rates chan model.ExchangeRate, waitGroup *sync.WaitGroup) {
	rates <- PrivatFetcher{}.Fetch()
	waitGroup.Done()
}

func FetchLatestQuote(quotes chan model.CryptoQuote, waitGroup *sync.WaitGroup) {
	quotes <- CmcFetcher{}.Fetch()
	waitGroup.Done()
}

func OneTimeFetch() string {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	var rates = make(chan model.ExchangeRate)
	var quotes = make(chan model.CryptoQuote)

	go FetchExchangeRate(rates, &waitGroup)
	go FetchLatestQuote(quotes, &waitGroup)

	exchangeRateUSD := <-rates
	quote := <-quotes

	rateUAHtoUSD, err := strconv.ParseFloat(exchangeRateUSD.Sale, 64)
	if err != nil {
		log.Print(err)
	}

	var cryptos []model.CryptoVerbose

	cryptos = append(cryptos, model.CryptoVerbose{
		Name:             quote.Data.BTC.Symbol,
		PriceInUAH:       quote.Data.BTC.Quote.USD.Price * rateUAHtoUSD,
		PriceInUSD:       quote.Data.BTC.Quote.USD.Price,
		MaxSupply:        quote.Data.BTC.MaxSupply,
		TotalSupply:      quote.Data.BTC.TotalSupply,
		PercentChange1H:  quote.Data.BTC.Quote.USD.PercentChange1H,
		PercentChange24H: quote.Data.BTC.Quote.USD.PercentChange24H,
		PercentChange7D:  quote.Data.BTC.Quote.USD.PercentChange7D,
		PercentChange30D: quote.Data.BTC.Quote.USD.PercentChange30D,
		MarketCap:        quote.Data.BTC.Quote.USD.MarketCap,
	})

	cryptos = append(cryptos, model.CryptoVerbose{
		Name:             quote.Data.DOGE.Symbol,
		PriceInUAH:       quote.Data.DOGE.Quote.USD.Price * rateUAHtoUSD,
		PriceInUSD:       quote.Data.DOGE.Quote.USD.Price,
		MaxSupply:        quote.Data.DOGE.MaxSupply,
		TotalSupply:      quote.Data.DOGE.TotalSupply,
		PercentChange1H:  quote.Data.DOGE.Quote.USD.PercentChange1H,
		PercentChange24H: quote.Data.DOGE.Quote.USD.PercentChange24H,
		PercentChange7D:  quote.Data.DOGE.Quote.USD.PercentChange7D,
		PercentChange30D: quote.Data.DOGE.Quote.USD.PercentChange30D,
		MarketCap:        quote.Data.DOGE.Quote.USD.MarketCap,
	})

	var message string

	for _, crypto := range cryptos {
		message += crypto.String()
	}

	currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
	message += fmt.Sprintf("<i>%v</i>\n", currentTime)

	return message
}
