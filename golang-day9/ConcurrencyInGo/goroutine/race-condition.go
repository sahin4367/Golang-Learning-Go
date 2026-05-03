package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	raceExample()
	raceExampleFixed()
	raceExampleFixedWithAtomatic()
}

func raceExampleFixedWithAtomatic() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var val int32 = 0

	go func() {
		for i := 0; i < 1000000; i++ {
			atomic.AddInt32(&val, 1)
			//Primitive tipler ucun atomatic -=>paketi var isitifade ede bilirik
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)

}

func raceExampleFixed() {
	//Bunu bilek ki biz Mutex den isitifade ederek : eyni anda tek tek bir isitifadeciye yetise bilir!
	wg := sync.WaitGroup{}
	wg.Add(2)

	lock := sync.Mutex{} // burda bir deyisen yaradiriq lock kimi | Mutex tipinde~!

	val := 0 //shared
	go func() {
		for i := 0; i < 10000; i++ {
			lock.Lock()   // bloklayiirq
			val++         //deyeri artiriiriq
			lock.Unlock() // ve unlock ederek blokdan azad edirik
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			lock.Lock()   // bloklayiirq
			val++         //deyeri artiriiriq
			lock.Unlock() // ve unlock ederek blokdan azad edirik
		}

		wg.Done()
	}()

	wg.Wait()
	fmt.Println(val)
}

func raceExample() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	val := 0
	go func() {
		for i := 0; i < 1000000; i++ {
			val++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			val++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}
