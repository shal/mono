# Public API

Create new public API client.

```go
public := NewPublic()
```

Get currency rates.

```go
rates, err := public.Rates()
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

You can create custom requests:

* **POST** request using `public.PostJSON(...)` method.
* **GET** request using `public.GetJSON(...)` method.
