package fetchers

import (
	"../model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type PrivatFetcher struct {
}

func (pF PrivatFetcher) Fetch() model.ExchangeRate {
	resp, err := http.Get(os.Getenv("PRIVAT_URL"))
	if err != nil {
		log.Print(err)
	}

	var exchangeRates []model.ExchangeRate

	err = json.NewDecoder(resp.Body).Decode(&exchangeRates)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	rateUSD, err := getUSD(exchangeRates)
	if err != nil {
		log.Print(err)
	}

	return rateUSD
}

func getUSD(exchangeRates []model.ExchangeRate) (model.ExchangeRate, error) {
	for _, exchangeRate := range exchangeRates {
		if exchangeRate.Currency == "USD" {
			return exchangeRate, nil
		}
	}
	return model.ExchangeRate{}, errors.New("no USD exchange rate fetched")
}
