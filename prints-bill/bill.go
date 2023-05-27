package main

import (
	"fmt"
	"math"
)

type Play struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Plays map[string]Play

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func (play Play) playKind() string {
	return play.Kind
}

func (play Play) playName() string {
	return play.Name
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func (play Play) amountFor(audience int) float64 {
	amount := 0.0
	switch play.playKind() {
	case "tragedy":
		amount = 40000
		if audience > 30 {
			amount += 1000 * (float64(audience - 30))
		}
	case "comedy":
		amount = 30000
		if audience > 20 {
			amount += 10000 + 500*(float64(audience-20))
		}
		amount += 300 * float64(audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", play.playKind()))
	}
	return amount
}

func (play Play) volumeCreditsFor(audience int) float64 {
	credits := 0.0
	credits += math.Max(float64(audience-30), 0)
	if play.playKind() == "comedy" {
		credits += math.Floor(float64(audience / 5))
	}
	return credits
}

type Rates []Rate

func (rates Rates) totalAmount() float64 {
	amount := 0.0
	for _, rate := range rates {
		amount += rate.Amount
	}
	return amount
}

func (rates Rates) totalVolumeCredits() float64 {
	credits := 0.0
	for _, rate := range rates {
		credits += rate.Credits
	}
	return credits
}

type Player interface {
	playKind() string
	playName() string
	amountFor(audience int) float64
	volumeCreditsFor(audience int) float64
}

type Rate struct {
	Play     Player
	Amount   float64
	Audience int
	Credits  float64
}

type Bill struct {
	Customer           string
	TotalAmount        float64
	TotalVolumeCredits float64
	Rates              Rates
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, rate := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", rate.Play.playName(), rate.Amount/100, rate.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

// func renderHtml(bill Bill) string {
// 	result := fmt.Sprintf("<h1>Statement for %s</h1>\n", bill.Customer)
// 	result += "<table>\n"
// 	result += "<tr><th>play</th><th>seats</th><th>cost</th></tr>"
// 	for _, rate := range bill.Rates {
// 		result += fmt.Sprintf("<tr><td>%s</td><td>%d</td><td>$%.2f</td></tr>\n", rate.Play.playName(), rate.Audience, rate.Amount/100)
// 	}
// 	result += "</table>\n"
// 	result += fmt.Sprintf("<p>Amount owed is <em>$%.2f</em></p>\n", bill.TotalAmount/100)
// 	result += fmt.Sprintf("<p>you earned <em>%.0f</em> credits</p>\n", bill.TotalVolumeCredits)
// 	return result
// }

func statement(invoice Invoice, plays Plays) string {
	var rates Rates
	for _, perf := range invoice.Performances {
		play := playFor(plays, perf)
		audience := perf.Audience
		rates = append(rates, Rate{
			Play:     play,
			Amount:   play.amountFor(audience),
			Audience: audience,
			Credits:  play.volumeCreditsFor(audience),
		})
	}

	bill := Bill{
		Customer:           invoice.Customer,
		TotalAmount:        rates.totalAmount(),
		TotalVolumeCredits: rates.totalVolumeCredits(),
		Rates:              rates,
	}

	return renderPlainText(bill)
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := map[string]Play{
		"hamlet":  {Name: "Hamlet", Kind: "tragedy"},
		"as-like": {Name: "As You Like It", Kind: "comedy"},
		"othello": {Name: "Othello", Kind: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
