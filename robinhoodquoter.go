package rbh

import (
	"rbh/ep"
	"rbh/http2"
	"rbh/rh"
	"strconv"
	"time"
)

// rbh.APIQuoter
type APIQuoter struct {
	c *http2.Client
	s string // ticker e.g. TCKR
}

// q, err := rbh.NewAPIQuoter(timeout, ticker)
func NewAPIQuoter(timeout time.Duration, ticker string) (*APIQuoter, error) {
	// i
	q := &APIQuoter{&http2.Client{Timeout: http2.globalTimeout}, ticker}
	if _, err := q.Quote(); err != nil {
		return nil, err
	}
	// ii
	q.c.Timeout = timeout
	return q, nil
}

// price, err := q.Quote()
func (q *APIQuoter) Quote() (float64, error) {
	// i
	var quote rh.Quote2
	if err := q.c.Get(ep.Base+"/quotes/"+q.s+"/", quote); err != nil {
		return 0, err
	}
	// ii
	return quote.LastKnownPrice()
}
