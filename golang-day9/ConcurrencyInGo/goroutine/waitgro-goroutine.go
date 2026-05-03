package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("go funkisyaininin runi")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("hello main go")
}
