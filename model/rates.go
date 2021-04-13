package model

type ExchangeRate struct {
	Currency     string `json:"ccy,omitempty"`
	BaseCurrency string `json:"base_ccy,omitempty"`
	Buy          string `json:"buy,omitempty"`
	Sale         string `json:"sale,omitempty"`
}
