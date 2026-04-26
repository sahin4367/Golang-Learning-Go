package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	b.mu.Lock()
	time.Sleep(10 * time.Millisecond)
	b.balance += amount
	b.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	account := &BankAccount{balance: 100}

	// 1. WaitGroup istifadəsi: 10 nəfər eyni anda pul yatırır
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go account.Deposit(10, &wg)
	}

	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		fmt.Printf("Bütün əməliyyatlar bitdi. Son balans: %d AZN\n", account.balance)
	case <-time.After(1 * time.Second):
		fmt.Println("Xəta: Əməliyyat çox uzun çəkdi (Timeout)!")
	}
}
