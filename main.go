package main

import (
	"fmt"
	"log"
	"os"
	"rbh"
	"rbh/algo"
	"strings"
)

var symbol = os.Args[2]

func main() {
	// // i
	// acc, err := rbh.Login("username", "password")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer acc.Logout()
	// ii-iii
	quote, err := rbh.NewCSVQuoter(strings.Join(os.Args[1:3], "/"), ' ') // [2006-01-02 TCKR]
	if err != nil {
		log.Fatalln(err)
	}

	model := algo.NewModel(2000)
	// iv
	var open, market float64
	for i := 0; i < 23400; i++ { // t := range interval.New(mkt.Open; mkt.Close, time.Second) {
		// a
		price, err := quote.Quote()
		if err != nil {
			log.Println(err)
			continue
		}
		// fmt.Printf("%0.4f\n", price)

		// pretty
		if i == 0 {
			open = price
		} else if i == 23399 {
			market = price - open
		}
		// price := quote.LastTradePrice
		// b
		shares := model.Solve(price)
		if shares == 0 {
			continue
		}
		// c
		// order, err := acc.Limit(d-time.Since(t), symbol, shares, price, rbh.GFD) // remaining duration
		// if err != nil {
		// 	log.Println(err)
		// 	continue
		// }
		// orders = append(orders, order)
	}
	// v
	algorithm, orders := model.Stats()
	fmt.Printf("open: %6.2f market: %+6.2f %+6.2f%% algthm: %+6.2f %+6.2f%% orders: %3d\n", open, market, 100*market/open, algorithm, 100*algorithm/open, orders)
	//	fmt.Printf("%.2f\t%d\n", algorithm, orders)
}
