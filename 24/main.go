package main

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance(o Point) float64 {
	sqr := func(x float64) float64 { return x * x }
	return math.Sqrt(sqr(o.x-p.x) + sqr(o.y-p.y))
}
