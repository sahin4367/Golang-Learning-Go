package main

import (
	"fmt"
	"math"
)

// 1. Struct ilə Normal Function
type Numbers struct {
	X float64
	Y float64
}

func Abs(n Numbers) float64 {
	return math.Sqrt(n.X*n.X + n.Y*n.Y)
}

// 2. Custom Type ilə Function
type MyFloat float64

func AbsFloat(f MyFloat) float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 3. Method (Value Receiver)
type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 4. Method (Pointer Receiver)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 5. Function ilə Pointer istifadə
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// MAIN
func main2() {
	// 1️
	n := Numbers{6, 8}
	fmt.Println("Abs Numbers:", Abs(n))

	// 2️
	f := MyFloat(-math.Sqrt2)
	fmt.Println("Abs MyFloat:", AbsFloat(f))

	// 3️
	v := Vertex{3, 4}
	fmt.Println("Vertex Abs:", v.Abs())

	// 4️
	v.Scale(10)
	fmt.Println("Scaled Vertex:", v)

	// 5️
	p := &Vertex{4, 3}
	ScaleFunc(p, 2)
	fmt.Println("Pointer Function:", p)
}
