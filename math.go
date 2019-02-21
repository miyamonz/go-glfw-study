package main

import (
	"math"
)

func sqrt(v float32) float32 {
	r := math.Sqrt(float64(v))
	return float32(r)
}
func sin(v float32) float32 {
	r := math.Sin(float64(v))
	return float32(r)
}
func cos(v float32) float32 {
	r := math.Cos(float64(v))
	return float32(r)
}
func tan(v float32) float32 {
	r := math.Tan(float64(v))
	return float32(r)
}
