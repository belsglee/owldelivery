package main

import (
	"fmt"
	"lib"
	"os"
	"strings"
)

func main() {

	vset := os.Args[1]
	symb := os.Args[2]

	files := lib.Catchfile(vset, symb)
	fmt.Println(files)

	for _, file := range files {
		fmt.Println(file)
		fmt.Printf("%-5s\t%-10s\n", symb, strings.Split(file, "/")[2])
		lib.Inspect_para(file, 1000, 2000)
	}

}
