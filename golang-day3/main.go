package main

import (
	"fmt"
	"os"
)

// BASIC FUNCTIONS
func sumNum(a, b int32) int32 {
	return a + b
}

// Greeting function
func hiName(name string) string {
	return "Hi my friend, " + name
}

// Variadic function (unknown number of params)
func totalNum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// MULTIPLE RETURN
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Cannot divide by zero"
	}
	return a / b, "Success"
}

// DEFER / PANIC / RECOVER
func testRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover edildi:", r)
		}
	}()

	fmt.Println("1. Test başladı")

	panic("Nəsə səhv oldu")

	// Bu setir işləməyəcək
	// çünki panic atıldı
	fmt.Println("2. Bu setir işləməyəcək")
}

func readFile() {
	file, err := os.Open("ozun-oxu.txt")

	if err != nil {
		fmt.Println("Xəta:", err)
		return
	}

	defer file.Close()

	fmt.Println("File uğurla açıldı")
}

func main() {

	fmt.Println("Main start\n")

	// Basic functions
	fmt.Println("Sum:", sumNum(2, 5))
	fmt.Println(hiName("Shahin"))
	fmt.Println("Total:", totalNum(1, 2, 3, 4))

	fmt.Println("\n--- Multiple Return ---")
	result, message := divide(42, 6)
	fmt.Println("Result:", result, "|", message)

	fmt.Println("\n--- File Handling ---")
	readFile()

	fmt.Println("\n--- Panic / Recover ---")
	testRecover()

	fmt.Println("\n Main davam edir (crash olmadı)")
}
