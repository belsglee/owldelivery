package main

import (
	"fmt"
	"rbh/lib"
	"log"
	"rbh"
	"rbh/algo"
	"strings"
)

func main() {

	vset := "quotes/2018-03-28/"
	symbols := getSymbols()

	fmt.Println(vset, symbols)
	fmt.Println("got symbols")

	for _, symbol := range symbols {

		files := lib.Catchfile(vset, symbol)

		for _, file := range files {
			fmt.Printf("%-5s\t%-10s\t", symbol, strings.Split(file, "/")[1])
			main_sub(file)
		}
	}

	fmt.Println(len(symbols))

}

/*func uitext() string {

	var vset string

	fmt.Println("Hello! Welcome to train and valid service for Owl Delivery")

	fmt.Print("Validation set: ")
	fmt.Scanln(&vset)

	return vset

}*/

func getSymbols() []string {

	folder := "quotes/2018-03-28/"
	files := lib.List(folder)

	var symbols []string

	for _, file := range files {
		symbols = append(symbols, strings.Split(file, "/")[2])
	}

	return symbols

}


func main_sub(file string) {
	quoter, err := rbh.NewCSVQuoter(file, ' ') // [2006-01-02 TCKR]
	if err != nil {
		log.Fatalln(err)
	}

	model := algo.NewModel(100)

	var open float64
	for i := 0; i < 23400; i++ { // t := range interval.New(mkt.Open; mkt.Close, time.Second) {
		// a
		price, err := quoter.Quote()

		if err != nil {
		//	log.Println(err)
			continue
		}

		if i == 0 {
			open = price
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
	fmt.Printf("open: %6.2f algthm: %+6.2f %+6.2f%% #oforders: %3d\n", open, algorithm, 100*algorithm/open, orders)
}
