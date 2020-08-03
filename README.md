# MonoBank SDK

[godoc]: https://godoc.org/shal.dev/mono
[godoc-img]: https://godoc.org/shal.dev/mono?status.svg

[ci]: https://circleci.com/gh/shal/mono
[ci-img]: https://circleci.com/gh/shal/mono.svg?style=svg

[goreport]: https://goreportcard.com/report/shal.dev/mono
[goreport-img]: https://goreportcard.com/badge/shal.dev/mono

[version]: https://img.shields.io/github/v/tag/shal/mono?sort=semver

[![Circle CI][ci-img]][ci]
[![Docs][godoc-img]][godoc]
[![Go Report][goreport-img]][goreport]
[![Version][version]][version]

:bank: Golang client for [Mono API](https://api.monobank.ua/docs/).

![Monobank API](assets/logo.png)

You can find documentation for all types of API [here](./docs).

1. [Introduction](#introduction)
2. [Documentation](#documentation)
3. [Using library](#use)
4. [Example](#example)
5. [Contributions](#contributions)

## Introduction

Read access APIs for monobank app.

Please, use personal API only in non-commersial services.

If you have a service or application and you want to centrally join the API for customer service, you need to connect to a corporate API that has more features.

This will allow monobank clients to log in to your service (for example, in a financial manager) to provide information about the status of an account or statements.

## Documentation

As far as monobank have 3 types of API, we prepated three usage documentations:

* [Public](./docs/public.md)
* [Personal](./docs/personal.md)
* [Corporate](./docs/corporate.md)

## Use

This package has no dependencies, install it with command below

```sh
go get shal.dev/mono
```

You can take a look and inspire by following [examples](./examples)

## Example

```go
package main

import (
    "fmt"
    "os"
    "time"

    "shal.dev/mono"
)

func main() {
    personal := mono.NewPersonal("token")

    user, err := personal.User(context.Background())
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    from := time.Now().Add(-730 * time.Hour)
    to := time.Now()

    var account mono.Account

    for _, acc := range user.Accounts {
        ccy, _ := iso4217.CurrencyFromISO4217(acc.CurrencyCode)
        if ccy.Code == "UAH" {
            account = acc
        }
    }

    transactions, err := personal.Transactions(context.Background(), account.ID, from, to)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Printf("Account: %s\n", account.ID)

    fmt.Println("Transactions:")
    for _, transaction := range transactions {
        fmt.Printf("%d\t%s\n", transaction.Amount, transaction.Description)
    }
}
```

More about this example [here](./examples/personal/main.go).

![Example](./examples/personal/report.png)

## License

Project released under the terms of the MIT [license](./LICENSE).
