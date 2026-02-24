package main

import (
	"fmt"
	"reflect"
)

func main() {
	// VARIABLES
	var a int = 10
	var name string = "Shahin"
	var isActive bool = true
	var price float32 = 5.99

	x := 25

	fmt.Println("Variables:")
	fmt.Println(a, name, isActive, price, x)
	a = 50
	fmt.Println("Updated a:", a)

	// CONSTANTS
	const pi float64 = 3.14
	const appName = "Go Learning"

	fmt.Println("\nConstants:")
	fmt.Println(pi, appName)

	// DATA TYPES CHECK
	fmt.Println("\nType Checking:")
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of isActive: %T\n", isActive)

	t := reflect.TypeOf(price)
	fmt.Println("Type of price (reflect):", t)

	// TYPE CONVERSION
	var num int = 5
	var floatNum float64 = float64(num)
	var uintNum uint = uint(floatNum)

	fmt.Println("\nType Conversion:")
	fmt.Println("int:", num)
	fmt.Println("float64:", floatNum)
	fmt.Println("uint:", uintNum)

	// FINAL OUTPUT
	fmt.Println("\nProgram finished successfully 🚀")
}
