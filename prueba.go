package main

import (
	"fmt"
	"math"
)

type abser interface {
	abs()	float64
}

type myfloat float64

func (f myfloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex struct {
	x, y float64
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.x + v.y)
}

func main() {

	var a abser
	f := myfloat(-4)
	v := vertex{3,4}

	a = f
	fmt.Println(a.abs())
	a = &v

	fmt.Println(a.abs())

}