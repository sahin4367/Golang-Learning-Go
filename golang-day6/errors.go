package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it did not work",
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError{
			time.Now(),
			fmt.Sprintf("cannot Sqrt negative number: %v", x),
		}
	}
	return math.Sqrt(x), nil
}

func main2() {
	fmt.Println("--- Strconv Test ---")
	input := "43" // "43t" yazsan xəta bölməsi işləyəcək
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Conversions error : ", err)
	} else {
		fmt.Println("Convertions number : ", number)
	}

	fmt.Println("\n--- Custom Error Test ---")
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n--- Sqrt Error Test ---")
	val1, err1 := Sqrt(2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Sqrt(2):", val1)
	}

	val2, err2 := Sqrt(-2)
	if err2 != nil {
		fmt.Println(err2) // Burada bizim yazdığımız MyError işə düşəcək
	} else {
		fmt.Println("Sqrt(-2):", val2)
	}
}
