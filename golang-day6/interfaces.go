package main

import (
	"fmt"
	"math"
)

/*
GO INTERFACES :
|- Interface nədir?

Interface Go-da methodların toplusudur.

Qayda:
Əgər bir type interface-də olan bütün methodları
implement edirsə, o interface-i avtomatik implement etmiş sayılır.

Go-da:

implements
extends

kimi sözlər YOXDUR.

Bu sistem "Implicit Implementation" adlanır.
*/

type Number interface {
	Abs() float64
}

type MyFloatt float64

func (f MyFloatt) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Point struct {
	X float64
	Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Basic Interface Example

type Printer interface {
	Print()
}

type Message struct {
	Text string
}

func (m Message) Print() {
	fmt.Println(m.Text)
}

func main() {

	// Interface ilə işləmə

	var n Number

	f := MyFloatt(-math.Sqrt2)
	p := &Point{3, 4}

	n = f
	fmt.Println("MyFloat Abs:", n.Abs())

	n = p
	fmt.Println("Point Abs:", n.Abs())

	// Sadə interface example

	var pr Printer
	pr = Message{"Salam gözəl insan."}

	pr.Print()
}
