package lib

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Scandate() string {

	fmt.Println("Available date from 2017-06-26 to 2017-10-30")

	fmt.Print("Enter date: ")
	var input string
	fmt.Scanln(&input)

	return input
}

func Scansymbol() string {

	fmt.Print("Enter symbol name: ")
	var input string
	fmt.Scanln(&input)

	return input
}

func Catchfile(dirPath string, symbol string) []string {

	var list []string

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		matched, _ := filepath.Match(symbol, info.Name())
		if matched {
			list = append(list, path)
		}
		return nil
	})
	return list
}

func List(dirname string) []string {

	var list []string

	a, _ := ioutil.ReadDir(dirname)

	for _, l := range a {
		list = append(list, dirname+l.Name())
	}

	return list
}

func Readfile(name string) []float64 {

	f, err := os.Open(name)

	Handle(err)

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ' '

	records, err := r.ReadAll()
	Handle(err)


	var record []float64
	var price string

	for _, prices := range records {
		fmt.Println(prices[0])
		price = prices[0]
		record = append(record, to_float(price[:len(price)-2]))
	}


	return record
}

func MK(quote []float64) float64 {
	return quote[len(quote)-1] - quote[0]
}

func Handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func to_float(price string) float64 {

	p, _ := strconv.ParseFloat(price, 64)
	return p

}

/*
func Readfile2(name string) []int {

	f, err := os.Open(name)

	Handle(err)

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ' '

	records, err := r.ReadAll()
	Handle(err)

	var record []int

	for _, price := range records {
		p, _ := strconv.Atoi(strings.Replace(price[1], ".", "", 1))
		record = append(record, p)
	}
	return record
} */

/*func double_Lazy_linear_reg(price [5]float64) float64 {

	var m float64

	for i := 0; i < 5; i++ {
		m = m + ((5*(float64(i)+1) - 15) * price[i])
	}

	return m
}*/
