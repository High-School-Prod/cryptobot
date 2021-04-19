package logic

import (
	"github.com/NazarNintendo/cryptobot/model"
	"log"
	"strconv"
)

func formatCrypto(exchangeRateUSD model.ExchangeRate, currency model.Currency) model.Crypto {
	priceInUSD := currency.Quote.USD.Price

	rateUAHtoUSD, err := strconv.ParseFloat(exchangeRateUSD.Sale, 64)
	if err != nil {
		log.Print(err)
	}

	priceInUAH := priceInUSD * rateUAHtoUSD

	crypto := model.Crypto{
		Name:             currency.Symbol,
		PriceInUAH:       priceInUAH,
		PriceInUSD:       priceInUSD,
		PercentChange24H: currency.Quote.USD.PercentChange24H,
	}

	//baseAggregator := BaseAggregator{aggregationCounter, historicalQuotes}

	//Aggregate(Aggregator10M{baseAggregator}, &crypto)
	//Aggregate(Aggregator1H{baseAggregator}, &crypto)
	//Aggregate(Aggregator2H{baseAggregator}, &crypto)

	return crypto
}
