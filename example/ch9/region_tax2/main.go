package main

import (
	"fmt"
	"log"
	"os"

	"github.com/learning-go-book/simpletax/v2"
	"github.com/shopspring/decimal"
)

func main() {
	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	zip := os.Args[2]
	country := os.Args[3]
	parsent, err := simpletax.ForCountryPostalCode(country, zip)
	if err != nil {
		log.Fatal(err)
	}
	total := amount.Add(amount.Mul(parsent)).Round(2)
	fmt.Println(total)
}
