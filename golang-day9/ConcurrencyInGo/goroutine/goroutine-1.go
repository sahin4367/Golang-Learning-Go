package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Qirmizi Isiq~!")
	}()

	fmt.Println("Yasil  Isiq~!")
}
