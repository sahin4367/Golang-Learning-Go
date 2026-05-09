package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("oxu.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// This  is logic --> linebyline methosd scanners !!!
