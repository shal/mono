# Public API

Create new public API client.

```go
public := mono.NewPublic()
```

Get currency rates.

```go
rates, err := public.Rates(context.Background())
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

for _, rate := range rates {
    ccyA, _ := iso4217.CurrencyFromISO4217(rate.CodeA)
    ccyB, _ := iso4217.CurrencyFromISO4217(rate.CodeB)

    if rate.RateBuy != 0 {
        fmt.Printf("%s/%s - %f\n", ccyA.Name, ccyB.Name, rate.RateBuy)
    } else {
        fmt.Printf("%s/%s - %f\n", ccyA.Name, ccyB.Name, rate.RateCross)
    }
}
```

You can create custom requests:

* **POST** request using `public.PostJSON(...)` method.
* **GET** request using `public.GetJSON(...)` method.
