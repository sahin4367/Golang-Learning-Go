package main

import (
	"fmt"
	"math"
)

/*
GO INTERFACES :
- Interface nədir? Interface Go-da methodların toplusudur.
- Qayda: Əgər bir type interface-də olan bütün methodları implement edirsə, o interface-i avtomatik implement etmiş sayılır.
- Go-da `implements` və `extends` kimi sözlər YOXDUR. Bu sistem "Implicit Implementation" adlanır.
*/

// ------------------ Example 1: Basic Number Interface ------------------

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

// ------------------ Example 2: Printer Interface ------------------

type Printer interface {
	Print()
}

type Message struct {
	Text string
}

func (m Message) Print() {
	fmt.Println(m.Text)
}

// ------------------ Example 3: Nil Receiver ------------------

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("nil qaytarsinn!")
		return
	}
	fmt.Println(t.S)
}

// describe olaraq düzəldildi (orfoqrafiya)
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ------------------ Example 4: Type Assertion ------------------

func typeAssertionExamples() {
	var i interface{} = 3

	// Tip eynidirsə, value qaytarır
	// Diqqət: Əgər i int olmasaydı, bu sətir proqramı çökdürərdi (panic).
	// Ona görə həmişə "ok" ilə yoxlamaq yaxşıdır.
	tip, ok := i.(int)
	if ok {
		fmt.Println("Tip int-dir:", tip)
	}

	// İki dəyər ilə yoxlama (ok pattern)
	a, b := i.(string)
	fmt.Println("Value:", a) // value boş olacaq (""), çünki type fərqlidir
	fmt.Println("OK?:", b)   // false olacaq
}

// ------------------ Example 5: Type Switch ------------------

func TipYoxla(i interface{}) {
	switch a := i.(type) {
	case int:
		fmt.Println("Integer:", a+2)
	case string:
		fmt.Println("String:", a+" salam.")
	case bool:
		fmt.Println("Bool:", a)
	default:
		fmt.Printf("Tipi bilinmir! Tip: %T\n", a)
	}
}

// ------------------ Example 6: Stringer Interface ------------------

type Person struct {
	name string
	age  int
}

// String methodu avtomatik fmt.Stringer interface-ni implement edir
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

// ------------------ MAIN ------------------

func main() {
	// Example 1: Number interface
	var n Number

	f := MyFloatt(-math.Sqrt2)
	p := &Point{3, 4}

	n = f
	fmt.Println("MyFloat Abs:", n.Abs())

	n = p
	fmt.Println("Point Abs:", n.Abs())

	// Example 2: Printer interface
	var pr Printer
	pr = Message{"Salam gözəl insan."}
	pr.Print()

	// Example 3: Nil receiver
	var i I
	var t *T

	i = t
	describe(i) // i-nin özü nil deyil (interface daxilində nil T tipi saxlayır)
	i.M()

	i = &T{"Salamun Aleykum"}
	describe(i)
	i.M()

	// Example 4: Type assertion
	typeAssertionExamples()

	// Example 5: Type switch
	TipYoxla(10)
	TipYoxla("Shahin")
	TipYoxla(false)
	TipYoxla(3.14)

	// Example 6: Stringer interface
	a := Person{"Vusal Quli", 23}
	b := Person{"Tural Boli", 18}
	// fmt.Println avtomatik olaraq String() metodunu çağırır
	fmt.Println(a, b)
}
