# Monobank API Client

:bank: Golang client for [Monobank API](https://api.monobank.ua/docs/).

## Install

This package has no dependencies, so just install it using command below

```sh
go get github.com/shal/mono
```

## Use

To create new API client:

```go
// Replace this string with your token.
c := mono.New("My token")
```

## Example

```go
import "fmt"

import "github.com/shal/mono"

func main() {
    token := "My token"

    client := mono.New(token)
    rates, err := client.Rates()
    if err != nil {
        panic(err)
    }

    for _, rate := range rates {
        fmt.Println(rate)
    }
}
```

## Contributions

You can send me some tips to [MonoBank](https://send.monobank.com.ua/2FVYpRHoi), if this package was useful.
