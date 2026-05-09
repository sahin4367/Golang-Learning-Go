package main

import (
	"io"
	"os"
)

func main() {
	source, _ := os.Open("oxu.txt")
	defer source.Close()

	create, _ := os.Create("yaz.txt")
	defer create.Close()

	io.Copy(create, source)
}
