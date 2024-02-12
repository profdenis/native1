package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{X: x, Y: y}
}

// String convertir le Point en String pour implémenter l'interface Stringer
func (p *Point) String() string {
	return fmt.Sprintf("Point(%f, %f)", p.X, p.Y)
}

// Add additionne 2 points et retourne le résultat dans un 3e point
func (p *Point) Add(other *Point) *Point {
	return &Point{p.X + other.X, p.Y + other.Y}
}

// Subtract soustrait 2 point et retourne le résultat dans un 3e point
func (p *Point) Subtract(other *Point) *Point {
	return &Point{p.X - other.X, p.Y - other.Y}
}

// Distance calcule et retourne la distance Euclidienne entre 2 point
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.X-p.X, 2) + math.Pow(other.Y-p.Y, 2))
}

// Scale mise à l'échelle d'un Point par un facteur et retourne le résultat dans un nouveau point
func (p *Point) Scale(factor float64) *Point {
	return &Point{p.X * factor, p.Y * factor}
}
