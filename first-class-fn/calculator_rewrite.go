package main

import (
	"fmt"
	"math"
)

type Calculator struct {
	acc float64
}

type opfunc func(float64, float64) float64

func (c *Calculator) Do(op opfunc, v float64) float64 {
	c.acc = op(c.acc, v)
	return c.acc
}

func Add(a, b float64) float64 { return a + b }

func Sub(a, b float64) float64 { return a - b }
func Mul(a, b float64) float64 { return a * b }

func main() {
	var c Calculator
	fmt.Println(c.Do(Add, 5)) // 5
	fmt.Println(c.Do(Sub, 3)) // 2
	fmt.Println(c.Do(Mul, 8)) // 16
	c.Do(Sqrt, 0)             // operand ignored

	// TO BE:
	// var c Calculator
	// c.Do(Add(1))
	// c.Do(Add(1))
	// c.Do(Sqrt())
	// c.Do(math.Sqrt) // 1.41421356237 c.Do(math.Cos) // 0.99969539804
}

func Sqrt(n, _ float64) float64 {
	return math.Sqrt(n)
}
