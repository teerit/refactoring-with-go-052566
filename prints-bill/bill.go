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

func playKind(play Play) string {
	return play.Kind
}

func playName(play Play) string {
	return play.Name
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func amountFor(plays Plays, perf Performance) float64 {
	amount := 0.0
	switch playKind(playFor(plays, perf)) {
	case "tragedy":
		amount = 40000
		if perf.Audience > 30 {
			amount += 1000 * (float64(perf.Audience - 30))
		}
	case "comedy":
		amount = 30000
		if perf.Audience > 20 {
			amount += 10000 + 500*(float64(perf.Audience-20))
		}
		amount += 300 * float64(perf.Audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", playKind(playFor(plays, perf))))
	}
	return amount
}

func volumeCreditsFor(plays Plays, perf Performance) float64 {
	credits := 0.0
	// add volume credits
	credits += math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == playKind(playFor(plays, perf)) {
		credits += math.Floor(float64(perf.Audience / 5))
	}
	return credits
}

func totalAmount(plays Plays, inv Invoice) float64 {
	amount := 0.0
	for _, perf := range inv.Performances {
		amount += amountFor(plays, perf)
	}
	return amount
}

func totalVolumeCredits(plays Plays, inv Invoice) float64 {
	credits := 0.0
	for _, perf := range inv.Performances {
		credits += volumeCreditsFor(plays, perf)
	}
	return credits
}

type Rate struct {
	Play     Play
	Amount   float64
	Audience int
}

type Bill struct {
	Customer           string
	TotalAmount        float64
	TotalVolumeCredits float64
	Rates              []Rate
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, rate := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", rate.Play.Name, rate.Amount/100, rate.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

func renderHtml(bill Bill) string {
	result := fmt.Sprintf("<h1>Statement for %s</h1>\n", bill.Customer)
	result += "<table>\n"
	result += "<tr><th>play</th><th>seats</th><th>cost</th></tr>"
	for _, rate := range bill.Rates {
		result += fmt.Sprintf("<tr><td>%s</td><td>%d</td><td>$%.2f</td></tr>\n", rate.Play.Name, rate.Audience, rate.Amount/100)
	}
	result += "</table>\n"
	result += fmt.Sprintf("<p>Amount owed is <em>$%.2f</em></p>\n", bill.TotalAmount/100)
	result += fmt.Sprintf("<p>you earned <em>%.0f</em> credits</p>\n", bill.TotalVolumeCredits)
	return result
}

func statement(invoice Invoice, plays Plays) string {
	var rates []Rate
	for _, perf := range invoice.Performances {
		rates = append(rates, Rate{
			Play:     playFor(plays, perf),
			Amount:   amountFor(plays, perf),
			Audience: perf.Audience,
		})
	}

	bill := Bill{
		Customer:           invoice.Customer,
		TotalAmount:        totalAmount(plays, invoice),
		TotalVolumeCredits: totalVolumeCredits(plays, invoice),
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
