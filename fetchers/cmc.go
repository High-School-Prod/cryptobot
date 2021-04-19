package fetchers

import (
	"encoding/json"
	"github.com/NazarNintendo/cryptobot/model"
	"log"
	"net/http"
	"os"
)

type CmcFetcher struct {
}

func (cF CmcFetcher) Fetch() model.CryptoQuote {
	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("CMC_URL"), nil)
	if err != nil {
		log.Print(err)
	}

	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("X-CMC_PRO_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	var historicalQuote model.CryptoQuote

	err = json.NewDecoder(resp.Body).Decode(&historicalQuote)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	return historicalQuote
}
