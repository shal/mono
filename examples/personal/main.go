package main

import (
	"fmt"
	"os"

	"github.com/shal/mono"
)

func main() {
	personal := mono.NewPersonal("token")

	user, err := personal.User()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("User: %s\n", user.Name)

	fmt.Println("Accounts:")
	for _, acc := range user.Accounts {
		ccy, _ := mono.CurrencyFromISO4217(acc.CurrencyCode)
		balance := fmt.Sprintf("%d.%d", acc.Balance/100, acc.Balance%100)
		fmt.Printf("%s - %s %s\n", ccy.Name, balance, ccy.Symbol)
	}
}
