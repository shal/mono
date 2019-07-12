# MonoBank SDK

[godoc]: https://godoc.org/github.com/shal/mono
[godoc-img]: https://godoc.org/github.com/shal/mono?status.svg

[ci]: https://circleci.com/gh/shal/mono
[ci-img]: https://circleci.com/gh/shal/mono.svg?style=svg

[goreport]: https://goreportcard.com/report/github.com/shal/mono
[goreport-img]: https://goreportcard.com/badge/github.com/shal/mono

[![Circle CI][ci-img]][ci]
[![Docs][godoc-img]][godoc]
[![Go Report][goreport-img]][goreport]

:bank: Golang client for [Mono API](https://api.monobank.ua/docs/).

## Install

This package has no dependencies, install it with command below

```sh
go get github.com/shal/mono
```

## Use

To create new API client:

```go
// Replace this string with your token.
client := mono.New("My token")
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
