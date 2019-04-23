package rbh

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

// rbh.CSVQuoter
type CSVQuoter struct{ r *csv.Reader }

// q, err := rbh.NewCSVQuoter(name, comma)
func NewCSVQuoter(name string, comma rune) (*CSVQuoter, error) {
	// i
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	// ii
	r := csv.NewReader(bufio.NewReader(file))

	// iii

	return &CSVQuoter{r}, nil
}

// price, err := q.Quote()
func (q *CSVQuoter) Quote() (float64, error) {
	// i
	record, err := q.r.Read()

	if err != nil {
		return 0, err
	}
	// ii
	price := record[1]

	return strconv.ParseFloat(price[:len(price)-2], 64)
}
