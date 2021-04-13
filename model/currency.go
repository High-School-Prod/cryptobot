package model

type CryptoQuote struct {
	Status Status `json:"status,omitempty"`
	Data   Data   `json:"data,omitempty"`
}

type Status struct {
	Timestamp    string `json:"timestamp,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	Elapsed      int    `json:"elapsed,omitempty"`
	CreditCount  int    `json:"credit_count,omitempty"`
	Notice       string `json:"notice,omitempty"`
}

type Data struct {
	DOGE Currency `json:"74,omitempty"`
	BTC  Currency `json:"1,omitempty"`
}

type Currency struct {
	Id                int      `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Symbol            string   `json:"symbol,omitempty"`
	Slug              string   `json:"slug,omitempty"`
	NumMarketPairs    int      `json:"num_market_pairs,omitempty"`
	DateAdded         string   `json:"date_added,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	MaxSupply         int      `json:"max_supply,omitempty"`
	CirculatingSupply float32  `json:"circulating_supply,omitempty"`
	TotalSupply       float64  `json:"total_supply,omitempty"`
	IsActive          int      `json:"is_active,omitempty"`
	Platform          string   `json:"platform,omitempty"`
	CmcRank           int      `json:"cmc_rank,omitempty"`
	IsFiat            int      `json:"is_fiat,omitempty"`
	LastUpdated       string   `json:"last_updated,omitempty"`
	Quote             Quote    `json:"quote,omitempty"`
}

type Quote struct {
	USD USD `json:"USD,omitempty"`
}

type USD struct {
	Price            float64 `json:"price,omitempty"`
	Volume24H        float64 `json:"volume_24h,omitempty"`
	PercentChange1H  float64 `json:"percent_change_1h,omitempty"`
	PercentChange24H float64 `json:"percent_change_24h,omitempty"`
	PercentChange7D  float64 `json:"percent_change_7d,omitempty"`
	PercentChange30D float64 `json:"percent_change_30d,omitempty"`
	PercentChange60D float64 `json:"percent_change_60d,omitempty"`
	PercentChange90D float64 `json:"percent_change_90d,omitempty"`
	MarketCap        float64 `json:"market_cap,omitempty"`
	LastUpdated      string  `json:"last_updated,omitempty"`
}
