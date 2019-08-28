package main

import (
	"fmt"
	"os"

	"github.com/shal/mono"
)

func main() {
	public := mono.NewPublic()

	rates, err := public.Rates()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, rate := range rates {
		ccyA, _ := mono.CurrencyFromISO4217(rate.CodeA)
		ccyB, _ := mono.CurrencyFromISO4217(rate.CodeB)

		if rate.RateBuy != 0 {
			fmt.Printf("%s/%s - %f\n", ccyA.Name, ccyB.Name, rate.RateBuy)
		} else {
			fmt.Printf("%s/%s - %f\n", ccyA.Name, ccyB.Name, rate.RateCross)
		}
	}
}
