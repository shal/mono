package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"github.com/shal/mono"
)

type Day struct {
	Expense float64
	Revenue float64
}

func transactions(token string) []mono.Transaction {
	personal := mono.NewPersonal(token)

	// Get information about current user.
	user, err := personal.User()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Find UAH account.
	account := mono.Account{}

	for _, acc := range user.Accounts {
		ccy, _ := mono.CurrencyFromISO4217(acc.CurrencyCode)
		if ccy.Code == "UAH" {
			account = acc
		}
	}

	// List all transactions for last month.
	transactions, err := personal.Transactions(account.ID, time.Now().Add(-730*time.Hour), time.Now())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return transactions
}

func main() {
	transactions := transactions("token")

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.X.Label.Text = "Time"
	p.Y.Label.Text = "UAH"

	days := make([]Day, 35)

	for _, t := range transactions {
		res := time.Unix(int64(t.Time), 0)
		if t.Amount < 0 {
			days[res.Day()].Expense += float64(-t.Amount/100) + float64(-t.Amount%100)
		} else {
			days[res.Day()].Revenue += float64(t.Amount/100) + float64(t.Amount%100)
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
		x += 1
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
