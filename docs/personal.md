# Personal API

Create new personal API client.

```go
// For more information about token: https://api.monobank.ua/.
personal := NewPersonal("token")
```

As far as public endpoints are also available, you can get currency rates.

```go
rates, err := personal.Rates()
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

for _, rate := range rates {
    ccyA := mono.CurrencyFromISO4217(rate.codeA)
    ccyB := mono.CurrencyFromISO4217(rate.codeB)

    fmt.Printf("%s/%s - %d", ccyA.Name, ccyB.Name, rate.RateSell)
}
```

Information about current user.

```go
user, err := personal.User()
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

fmt.Printf("User: %s\n", user.Name)

fmt.Println("Accounts:")
for _, acc := range user.Accounts {
    fmt.Println(acc)
}
```

User's transaction for a given period of time.

```go
from := time.Now().Add(-24 * time.Hour)
to := time.Now()

transactions, err := personal.Transactions(user.Accounts[0].ID, from, to)
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

fmt.Printf("Account: %s\n", user.Accounts[0].ID)

fmt.Println("Transactions:")
for _, transaction := range transactions {
    fmt.Printf("%d - %s\n",transaction.Amount, transcation.Description)
}
```

Set WebHook for give URI.

```go
_, err := personal.SetWebHook("http://example.com")
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
} else {
    fmt.Println("Success!")
}
```

You can create custom requests:

* **POST** request using `personal.PostJSON(...)` method.
* **GET** request using `personal.GetJSON(...)` method.
