package main

import (
	"context"
	"image/color"
	"log"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"shal.dev/mono"
	"shal.dev/mono/auth"
	"shal.dev/mono/iso4217"
)

type day struct {
	Expense float64
	Revenue float64
}

func transactions(ctx context.Context, token string) ([]mono.Transaction, error) {
	personal := auth.NewPersonal(token)

	client, err := mono.NewClient(personal)
	if err != nil {
		return nil, err
	}

	// Get information about current user.
	user, err := client.User(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Find UAH account.
	var account mono.Account

	for _, acc := range user.Accounts {
		if acc.Type == mono.Platinum || acc.Type == mono.Black {
			ccy, _ := iso4217.CurrencyFromISO4217(acc.CurrencyCode)
			if ccy.Code == "UAH" {
				account = acc
			}
		}
	}

	// List all transactions for last month.
	transactions, err := client.Transactions(ctx, account.ID, time.Now().Add(-730*time.Hour), time.Now(), nil)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token := os.Getenv("MONO_TOKEN")
	if token == "" {
		log.Fatal("MONO_TOKEN is not set")
	}

	transactions, err := transactions(ctx, token)
	if err != nil {
		log.Fatal(err)
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.X.Label.Text = "Time"
	p.Y.Label.Text = "UAH"

	days := make([]day, 35)

	for _, t := range transactions {
		if t.Amount < 0 {
			days[t.Time.Day()].Expense += float64(-t.Amount/100) + float64(-t.Amount%100)
		} else {
			days[t.Time.Day()].Revenue += float64(t.Amount/100) + float64(t.Amount%100)
		}
	}

	expenses := make(plotter.XYs, len(days))
	revenues := make(plotter.XYs, len(days))

	x := 0
	for _, v := range days {
		expenses[x].X = float64(x)
		revenues[x].X = float64(x)
		expenses[x].Y = v.Expense
		revenues[x].Y = v.Revenue
		x++
	}

	expensesPlot, err := plotter.NewLine(expenses)
	if err != nil {
		panic(err)
	}

	revenuesPlot, err := plotter.NewLine(revenues)
	if err != nil {
		panic(err)
	}

	expensesPlot.LineStyle.Color = color.RGBA{R: 255, A: 255}
	revenuesPlot.LineStyle.Color = color.RGBA{G: 100, A: 255}

	expensesCircles, err := plotter.NewScatter(expenses)
	if err != nil {
		panic(err)
	}

	revenuesCircles, err := plotter.NewScatter(revenues)
	if err != nil {
		panic(err)
	}

	expensesPlot.LineStyle.Color = color.RGBA{R: 255, A: 255}
	revenuesPlot.LineStyle.Color = color.RGBA{G: 100, A: 255}
	expensesCircles.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	revenuesCircles.GlyphStyle.Color = color.RGBA{G: 100, A: 255}

	expensesCircles.GlyphStyle.Shape = draw.CircleGlyph{}
	revenuesCircles.GlyphStyle.Shape = draw.CircleGlyph{}

	p.Add(expensesPlot, revenuesPlot, expensesCircles, revenuesCircles)

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 5*vg.Inch, "report.png"); err != nil {
		panic(err)
	}
}
