package main

import (
	"fmt"
	"time"
)

func main() {
	var i int // Dövrün xaricində bir dəfə elan olunur
	for i = 0; i < 10; i++ {
		go func() {
			fmt.Println(i) // Hamısı eyni 'i' referansına baxacaq
		}()
	}
	time.Sleep(time.Second * 3)
}
