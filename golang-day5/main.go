package main

import "fmt"

// ARRAYS

func arrayBasics() {

	// Array declaration
	var a [4]int
	a[0] = 10
	a[1] = 20

	fmt.Println("Array a:", a)

	// Initialize with values
	nums := [...]int{1, 2, 3, 4, 5}

	fmt.Println("Nums:", nums)
	fmt.Println("Length:", len(nums))
}

// ARRAY MEMORY (Pointer understanding)
func arrayMemory() {

	nums := [3]int32{12, 23, 4}

	fmt.Println("Address of nums[0]:", &nums[0])
	fmt.Println("Address of nums[1]:", &nums[1])
	fmt.Println("Address of nums[2]:", &nums[2])

}

// ARRAY COPY vs POINTER
func copyArray(a [3]int) {
	a[0] = 100
}

func pointerArray(a *[3]int) {
	a[0] = 100
}

func arrayCopyExample() {

	nums := [3]int{1, 2, 3}

	copyArray(nums)
	fmt.Println("After copyArray:", nums)

	pointerArray(&nums)
	fmt.Println("After pointerArray:", nums)
}

// MULTI-DIMENSION ARRAY
func multiArray() {

	var table [2][3]int32

	table[0] = [3]int32{1, 2, 3}
	table[1] = [3]int32{4, 5, 6}

	fmt.Println("Table:", table)
	fmt.Println("Element [1][2]:", table[1][2])
}

// SLICES
func sliceBasics() {

	s := []int{1, 2, 3}

	fmt.Println("Slice:", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}

// MAKE (Slice creation)
func sliceMake() {

	d := make([]int, 4)
	t := make([]int, 3, 5)

	t[2] = 5

	fmt.Println("Slice d:", d)
	fmt.Println("Slice t:", t)

	fmt.Println("Len:", len(t))
	fmt.Println("Cap:", cap(t))
}

// APPEND
func sliceAppend() {

	s := []int{1, 2, 3}

	s = append(s, 4)

	fmt.Println("After append:", s)
}

// MATRIX (Slice of slices)
func sliceMatrix() {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Last row:", matrix[len(matrix)-1])
	fmt.Println("Total rows:", len(matrix))
}

// MAPS
func mapExample() {

	ages := map[string]int{
		"Shahin": 22,
		"Ali":    30,
	}

	ages["Fuad"] = 25

	fmt.Println("Map:", ages)

	for name, age := range ages {
		fmt.Println(name, "->", age)
	}
}

func main() {

	fmt.Println("---- Arrays ----")
	arrayBasics()

	fmt.Println("\n---- Array Memory ----")
	arrayMemory()

	fmt.Println("\n---- Copy vs Pointer ----")
	arrayCopyExample()

	fmt.Println("\n---- Multi Array ----")
	multiArray()

	fmt.Println("\n---- Slice Basics ----")
	sliceBasics()

	fmt.Println("\n---- Slice Make ----")
	sliceMake()

	fmt.Println("\n---- Slice Append ----")
	sliceAppend()

	fmt.Println("\n---- Slice Matrix ----")
	sliceMatrix()

	fmt.Println("\n---- Maps ----")
	mapExample()

}
