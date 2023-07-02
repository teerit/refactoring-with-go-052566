package main

import (
	"fmt"
	"math"
)

type Calculator struct {
	acc float64
}

type opfunc func(v float64) float64

func (c *Calculator) Do(op opfunc) float64 {
	c.acc = op(c.acc)
	return c.acc
}

// func Add(a, b float64) float64 { return a + b }
// func Sub(a, b float64) float64 { return a - b }

func Mul(v float64) opfunc {
	return func(acc float64) float64 {
		return v * acc
	}
}

func Add(v float64) opfunc {
	return func(acc float64) float64 {
		return v + acc
	}
}

func Sub(v float64) opfunc {
	return func(acc float64) float64 {
		return acc - v
	}
}

func Div(v float64) opfunc {
	return func(acc float64) float64 {
		return acc / v
	}
}

func Sqrt() opfunc {
	return func(acc float64) float64 {
		return math.Sqrt(acc)
	}
}

func main() {
	// fmt.Println(c.Do(Add, 5)) // 5
	// fmt.Println(c.Do(Sub, 3)) // 2
	// fmt.Println(c.Do(Mul, 8)) // 16
	// c.Do(Sqrt, 0)             // operand ignored

	// TO BE:
	var c Calculator
	fmt.Println(c.Do(Add(5)))
	fmt.Println(c.Do(Sub(3)))
	fmt.Println(c.Do(Mul(8)))
	fmt.Println(c.Do(Sqrt()))
	c.Do(math.Sqrt) // 1.41421356237 c.Do(math.Cos) // 0.99969539804
}
