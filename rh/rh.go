package rh

import "time"

// rh.Quote
type Quote struct {
	AskPrice                    string    `json:"ask_price"`
	AskSize                     int       `json:"ask_size"`
	BidPrice                    string    `json:"bid_price"`
	BidSize                     int       `json:"bid_size"`
	LastTradePrice              string    `json:"last_trade_price"`
	LastExtendedHoursTradePrice string    `json:"last_extended_hours_trade_price"`
	PreviousClose               string    `json:"previous_close"`
	AdjustedPreviousClose       string    `json:"adjusted_previous_close"`
	PreviousCloseDate           string    `json:"previous_close_date"`
	Symbol                      string    `json:"symbol"`
	TradingHalted               bool      `json:"trading_halted"`
	HasTraded                   bool      `json:"has_traded"`
	LastTradePriceSource        string    `json:"last_trade_price_source"`
	UpdatedAt                   time.Time `json:"updated_at"`
	Instrument                  string    `json:"instrument"`
}

// rh.Quote2
type Quote2 struct {
	LastTradePrice              string `json:"last_trade_price"`
	LastExtendedHoursTradePrice string `json:"last_extended_hours_trade_price"`
}

// rh.Quotes
type Quotes struct {
	Results []Quote `json:"results"`
}
