package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("0905-2026.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
