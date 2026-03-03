package main

import "fmt"

// 🔹 Counter struct
type Counter struct {
	value int
}

// Pointer receiver (mutate edir)
func (c *Counter) Inc() {
	c.value++
}

// 🔹 Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Area (sadəcə oxuyur)
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Scale (dəyişir)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 🔹 swap function
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 🔹 slice example
func changeSlice(s []int) {
	s[0] = 999
}

func main() {

	// 1. Pointer + basic
	x := 10
	p := &x
	fmt.Println("Pointer value:", *p)

	// 2. Nil pointer (comment açma!)
	// var p2 *int
	// fmt.Println(*p2) // panic

	// 3. Counter (method receiver)
	c := Counter{value: 5}
	c.Inc()
	fmt.Println("Counter:", c.value) // 6

	// 4. Rectangle
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("Area:", rect.Area()) // 12

	rect.Scale(2)
	fmt.Println("New Area:", rect.Area()) // 48

	// 5. Swap
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println("Swap:", a, b) // 2 1

	// 6. Slice
	arr := []int{1, 2, 3}
	changeSlice(arr)
	fmt.Println("Slice:", arr) // [999 2 3]
}
