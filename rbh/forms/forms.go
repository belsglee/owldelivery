package forms

import (
	"net/url"
	"rbh/inspect"
	"strconv"
	"strings"
)

type Token struct {
	Username string `url:"username"`
	Password string `url:"password"`
}

// body := token.Encode()
func (t *Token) Encode() *strings.Reader { return encode(t) }

type Order struct {
	Account     string      `url:"account"`
	Trigger     string      `url:"trigger"`
	Type        string      `url:"type"`
	Symbol      string      `url:"symbol"`
	Instrument  string      `url:"instrument"`
	Side        string      `url:"side"`
	Quantity    int         `url:"quantity"`
	StopPrice   interface{} `url:"stop_price"`
	Price       float64     `url:"price"`
	TimeInForce string      `url:"time_in_force"`
}

// body := order.Encode()
func (o *Order) Encode() *strings.Reader { return encode(o) }

func encode(value interface{}) *strings.Reader {
	// i
	m := inspect.Struct(value, "url").Map()
	// ii
	var s, sep string
	for key, value := range m {
		s += sep + key + "="
		switch v := value.(type) {
		case string:  s += url.QueryEscape(v)
		case int:     s += strconv.Itoa(v)
		case float64: s += strconv.FormatFloat(v, 'f', 2, 64)
		}
		sep = "&"
	}
	// iii
	return strings.NewReader(s)
}
