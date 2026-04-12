package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		func(a int) {
			fmt.Println(a)
		}(i)
	}
	time.Sleep(time.Second * 3)
}
