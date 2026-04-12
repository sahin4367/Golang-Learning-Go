package main

import (
	"fmt"
	"sync"
	"time"
)

// ─── 1) Goroutine nümunəsi ───────────────────────────────────────────
func sayHello(name string) {
	fmt.Println("Salam:", name)
}

// ─── 2) WaitGroup nümunəsi ──────────────────────────────────────────
func runWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // hər goroutine üçün sayğacı artır
		go func(n int) {
			defer wg.Done() // goroutine bitəndə sayğacı azalt
			fmt.Println("WaitGroup goroutine:", n)
		}(i)
	}

	wg.Wait() // bütün goroutine-lər bitənə qədər gözlə
	fmt.Println("Hamısı tamamlandı!")
}

// ─── main ────────────────────────────────────────────────────────────
func main() {
	// --- Goroutine nümunəsi ---
	go sayHello("Oglan") // ayrı goroutine-də işləyir
	sayHello("Shahin")   // main goroutine-də işləyir
	time.Sleep(time.Second)

	fmt.Println("─────────────────")

	// --- WaitGroup nümunəsi ---
	runWaitGroup()
}
