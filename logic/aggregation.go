package logic

import (
	"github.com/NazarNintendo/cryptobot/model"
)

type Aggregator interface {
	aggregate(crypto *model.Crypto)
}

func Aggregate(aggregator Aggregator, crypto *model.Crypto) {
	aggregator.aggregate(crypto)
}

type Aggregator10M struct {
	BaseAggregator
}

type Aggregator1H struct {
	BaseAggregator
}

type Aggregator2H struct {
	BaseAggregator
}

type BaseAggregator struct {
	aggregationCounter int
	historicalQuotes   []model.CryptoQuote
}

func (a Aggregator10M) aggregate(crypto *model.Crypto) {
	var prevCounter int
	if a.aggregationCounter == 0 {
		prevCounter = 71
	} else {
		prevCounter = a.aggregationCounter - 1
	}
	newPrice := crypto.PriceInUSD
	oldPrice := getOldPrice(crypto.Name, a.historicalQuotes, prevCounter)
	crypto.PercentChange10M = getPercentChange(newPrice, oldPrice)
}

func (a Aggregator1H) aggregate(crypto *model.Crypto) {
	var prevCounter int
	if a.aggregationCounter < 6 {
		prevCounter = 71 - 5 + a.aggregationCounter
	} else {
		prevCounter = a.aggregationCounter - 6
	}
	newPrice := crypto.PriceInUSD
	oldPrice := getOldPrice(crypto.Name, a.historicalQuotes, prevCounter)
	crypto.PercentChange1H = getPercentChange(newPrice, oldPrice)
}

func (a Aggregator2H) aggregate(crypto *model.Crypto) {
	var prevCounter int
	if a.aggregationCounter < 12 {
		prevCounter = 71 - 11 + a.aggregationCounter
	} else {
		prevCounter = a.aggregationCounter - 12
	}
	newPrice := crypto.PriceInUSD
	oldPrice := getOldPrice(crypto.Name, a.historicalQuotes, prevCounter)
	crypto.PercentChange2H = getPercentChange(newPrice, oldPrice)
}

func getOldPrice(cryptoName string, oldQuotes []model.CryptoQuote, prevCounter int) float64 {
	var oldPrice float64
	if cryptoName == "DOGE" {
		oldPrice = oldQuotes[prevCounter].Data.DOGE.Quote.USD.Price
	} else {
		oldPrice = oldQuotes[prevCounter].Data.BTC.Quote.USD.Price
	}
	return oldPrice
}

func getPercentChange(newPrice float64, oldPrice float64) float64 {
	if oldPrice == 0. {
		return 0.
	} else {
		return (newPrice - oldPrice) / oldPrice * 100
	}
}
