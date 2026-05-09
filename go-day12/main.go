package main

import (
	"os"
)

func main() {
	data := []byte("Salam Shahin!")

	err := os.WriteFile("0905-2026.txt", data, 0644)

	if err != nil {
		panic(err)
	}
}
