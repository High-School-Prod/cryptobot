package model

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Client struct {
	User        *tb.User
	Monitorings []Monitoring
	Metrics     []Metric
}

type Monitoring struct {
	Currency  int
	Frequency int
}

type Metric struct {
	Currency  int
	Frequency int
	Growth    float64
}

func (c Client) String() string {
	return fmt.Sprintf("User = %v\nMonitorings = %v\nMetrics = %v", c.User.FirstName, c.Monitorings, c.Metrics)
}

func (m Monitoring) String() string {
	return fmt.Sprintf("Currency = %v\nFrequency = %v", m.Currency, m.Frequency)
}

func (m Metric) String() string {
	return fmt.Sprintf("Currency = %v\nFrequency = %v\nGrowth = %v", m.Currency, m.Frequency, m.Growth)
}

func (c *Client) AddMonitoring(currency int, frequency int) *Client {
	monitoring := Monitoring{Currency: currency, Frequency: frequency}
	for _, m := range c.Monitorings {
		if m == monitoring {
			return c
		}
	}
	c.Monitorings = append(c.Monitorings, monitoring)
	return c
}
