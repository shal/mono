# Corporate API

Create new corporate API client.

```go
// For more information about corporate API: https://api.monobank.ua/docs/corporate.html.
corporate := mono.NewCorporate(...)
```

As far as public endpoints are also available, you can get currency rates.

```go
rates, err := corporate.Rates(context.Background())
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
```

Information about authorized user by request ID.

```go
requestID := "xxx"
user, err := corporate.User(context.Background(), requestID)
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
```

User's transaction for a given period of time.

```go
requestID := "xxx"
user, err := corporate.User(context.Background(), requestID)
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

from := time.Now().Add(-730 * time.Hour)
to := time.Now()

var account mono.Account

for _, acc := range user.Accounts {
    ccy, _ := mono.CurrencyFromISO4217(acc.CurrencyCode)
    if ccy.Code == "UAH" {
        account = acc
    }
}

transactions, err := corporate.Transactions(context.Background(), requestID, account.ID, from, to)
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

fmt.Printf("Account: %s\n", account.ID)

fmt.Println("Transactions:")
for _, transaction := range transactions {
    fmt.Printf("%d\t%s\n", transaction.Amount, transaction.Description)
}
```

You can create custom requests:

* **POST** request using `corporate.PostJSON(...)` method.
* **GET** request using `corporate.GetJSON(...)` method.
