# MonoBank SDK

:bank: Golang client for [Mono API](https://api.monobank.ua/docs/).

![Logo](./assets/logo.png)

## Install

This package has no dependencies, install it with command below

```sh
go get github.com/shal/mono
```

## Use

To create new API client:

```go
// Replace this string with your token.
c := mono.New("My token")
```

TODO: Add usage of each API endpoint.

## Example

```go
package main

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
