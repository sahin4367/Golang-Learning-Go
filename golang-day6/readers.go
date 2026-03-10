package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	// Oxunan n sayda baytı (hərfi) tək-tək çeviririk
	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm') {
			p[i] += 13
		} else if (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return n, err
}

func main3() {
	fmt.Println("--- Sadə Reader Testi ---")
	r1 := strings.NewReader("Salamlar, müəllim!")
	d := make([]byte, 8)
	n, err := r1.Read(d)

	fmt.Printf("Oxunan bayt sayı: %d\n", n)
	fmt.Printf("Xəta (əgər varsa): %v\n", err)
	fmt.Printf("Bayt massivi: %v\n", d)
	fmt.Printf("Mətn formatında: %s\n", string(d)) // İlk 8 baytı göstərəcək

	fmt.Println("\n--- Rot13Reader Testi ---")

	// --- Sənin ikinci nümunən: Rot13 şifrəli Reader ---
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader{s}

	io.Copy(os.Stdout, &r2)
	fmt.Println() // Sonda yeni sətir
}
