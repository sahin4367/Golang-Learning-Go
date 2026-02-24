package main

import (
	"fmt"
)

func test() int {
	return 10
}

func main() {

	// IF / ELSE
	a := 3
	if a > 3 {
		fmt.Println("a 3-dən böyükdür")
	} else if a == 3 {
		fmt.Println("a 3-ə bərabərdir")
	} else {
		fmt.Println("a 3-dən kiçikdir")
	}

	// short if || kod daha suretli isleyir~!
	if b := 30; b < 50 {
		fmt.Println("b 50-dən kiçikdir:", b)
	}

	// SWITCH (single & multi case)
	score := 88
	switch score {
	case 56, 67:
		fmt.Println("Bu ballar 88 deyil")
	case 75, 85, 88:
		fmt.Println("Bu ballar 88-ə yaxındır")
	default:
		fmt.Println("Uyğun nəticə tapılmadı")
	}

	// type uyğunluğu
	//eger bele olsaydi value := "salam(yeni-string olsa CASE_lerde string olmalidir~!)"
	value := 3
	switch value {
	case 1:
		fmt.Println("value 1-dir")
	case 3:
		fmt.Println("value 3-dür")
	default:
		fmt.Println("value fərqlidir")
	}

	// function inside switch
	num := 10
	switch num {
	case test():
		fmt.Println("num 10-a bərabərdir")
	}

	// BASIC ERROR HANDLING
	fmt.Println("\nError Handling Example:")

	var x int = 0

	if x == 0 {
		fmt.Println("Error: x sıfır ola bilməz!")
	} else {
		fmt.Println("x:", x)
	}
}
