package main

import (
	"context"
	"fmt"
	"log"
	"shal.dev/mono/iso4217"
	"time"

	"shal.dev/mono"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	public, err := mono.NewPublic()
	if err != nil {
		log.Fatal(err)
	}

	rates, err := public.Rates(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, rate := range rates {
		ccyA, err := iso4217.CurrencyFromISO4217(rate.CodeA)
		if err != nil {
			log.Println(rate.CodeA)
			continue
		}
		ccyB, _ := iso4217.CurrencyFromISO4217(rate.CodeB)

		if rate.RateBuy != 0 {
			fmt.Printf("%s/%s - %f\n", ccyA.Name, ccyB.Name, rate.RateBuy)
		} else {
			fmt.Printf("%s/%s - %f\n", ccyA.Name, ccyB.Name, rate.RateCross)
		}
	}
}
