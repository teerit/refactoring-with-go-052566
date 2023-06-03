package main

import (
	"fmt"
	"math"
	"net/http"
)

type Calculator struct {
	acc float64
}

type opfunc func(float64) float64

func (c *Calculator) Do(fn opfunc) float64 {
	c.acc = fn(c.acc)
	return c.acc
}

func Add(a float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc + a
	}
}

func Sub(a float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc - a
	}
}

func Mul(a float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc * a
	}
}

func Sqrt() func(float64) float64 {
	return func(acc float64) float64 {
		return math.Sqrt(acc)
	}
}

func main() {
	// TO BE:
	var c Calculator
	c.Do(Add(8))
	c.Do(Sub(4))
	c.Do(Mul(4))
	c.Do(Sqrt())
	c.Do(math.Sqrt)
	// c.Do(math.Cos)

	fmt.Println("Accumulator:", c.acc)

	handler := func(w http.ResponseWriter, r *http.Request) {
	}

	http.HandleFunc("/", handler)
}

// func Sqrt(n, _ float64) float64 {
// 	return math.Sqrt(n)
// }
