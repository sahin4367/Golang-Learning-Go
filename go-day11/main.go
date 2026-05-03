package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CouponRequest struct {
	UserID int `json:"user_id"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var (
	jobs    = make(chan CouponRequest, 10)
	limiter = time.NewTicker(1 * time.Second) // Saniyədə 1 kupon (test üçün yavaşladılıb)
	wg      sync.WaitGroup
)

func main() {
	// 3 Worker işə salınır
	for w := 1; w <= 3; w++ {
		go worker(w)
	}

	http.HandleFunc("/get-coupon", handleCoupon)

	fmt.Println("Server 8080 portunda başladı...")
	http.ListenAndServe(":8080", nil)
}

func handleCoupon(w http.ResponseWriter, r *http.Request) {
	// CORS ayarları (Frontend üçün)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	req := CouponRequest{UserID: time.Now().Nanosecond()}

	select {
	case jobs <- req:
		json.NewEncoder(w).Encode(Response{Message: "Sorğu sıraya alındı!", Status: "success"})
	case <-ctx.Done():
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(Response{Message: "Sistem məşğuldur, limit aşılıb!", Status: "fail"})
	}
}

func worker(id int) {
	for req := range jobs {
		<-limiter.C // Throttling
		fmt.Printf("Worker %d: İstifadəçi %d üçün kupon verildi.\n", id, req.UserID)
		time.Sleep(500 * time.Millisecond)
	}
}
