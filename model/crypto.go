package model

import "fmt"

type Crypto struct {
	Name             string
	PriceInUAH       float64
	PriceInUSD       float64
	PercentChange10M float64
	PercentChange1H  float64
	PercentChange2H  float64
	PercentChange24H float64
}

func (c Crypto) String() string {
	return fmt.Sprintf("One <b>%v</b> is:\n\n", c.Name) +
		fmt.Sprintf("➤   <b>%.4f</b> UAH \n", c.PriceInUAH) +
		fmt.Sprintf("➤   <b>%.4f</b> USD\n\n", c.PriceInUSD) +
		fmt.Sprintf("Change in 10 minutes    ➝    <b>%.4v</b>%%\n", c.PercentChange10M) +
		fmt.Sprintf("Change in 1 hour             ➝    <b>%.4v</b>%%\n", c.PercentChange1H) +
		fmt.Sprintf("Change in 2 hours           ➝    <b>%.4v</b>%%\n", c.PercentChange2H) +
		fmt.Sprintf("Change in 24 hours         ➝    <b>%.4v</b>%%\n", c.PercentChange24H) +
		fmt.Sprintf("███████████████████████████████\n\n")
}

type CryptoVerbose struct {
	Name             string
	PriceInUAH       float64
	PriceInUSD       float64
	MaxSupply        int
	TotalSupply      float64
	PercentChange1H  float64
	PercentChange24H float64
	PercentChange7D  float64
	PercentChange30D float64
	MarketCap        float64
}

func (c CryptoVerbose) String() string {
	return fmt.Sprintf("One <b>%v</b> is:\n\n", c.Name) +
		fmt.Sprintf("➤   <b>%.4f</b> UAH \n", c.PriceInUAH) +
		fmt.Sprintf("➤   <b>%.4f</b> USD\n\n", c.PriceInUSD) +
		fmt.Sprintf("Max. supply   =   <b>%v</b> \n", c.MaxSupply) +
		fmt.Sprintf("Total supply   =   <b>%.4f</b> \n\n", c.TotalSupply) +
		fmt.Sprintf("Change in 1 hour           ➝    <b>%.4f</b>%%\n", c.PercentChange1H) +
		fmt.Sprintf("Change in 24 hours       ➝    <b>%.4f</b>%%\n", c.PercentChange24H) +
		fmt.Sprintf("Change in 7 days           ➝    <b>%.4f</b>%%\n", c.PercentChange7D) +
		fmt.Sprintf("Change in 30 days         ➝    <b>%.4f</b>%%\n\n", c.PercentChange30D) +
		fmt.Sprintf("Market capitilization   $<b>%.4f</b>\n", c.MarketCap) +
		fmt.Sprintf("███████████████████████████████\n\n")
}
