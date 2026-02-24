# 🚀 Introduction to Go (Golang)

This repository contains my first steps in learning the Go (Golang) programming language, including basic concepts, syntax, and examples.

---

## 📖 About Go

Go (Golang) is an open-source programming language developed by Google engineers — **Robert Griesemer, Rob Pike, and Ken Thompson**.

It was designed to be:

* Simple and readable
* Fast and efficient (compiled language)
* Great for backend and concurrent systems

---

## ⚙️ Setting up Environment

To start working with Go:

```bash
go version
```

This command checks whether Go is installed on your system.

---

## 📦 package main & fmt

Every Go program starts with:

```go
package main
```

This defines the entry point of the application.

### fmt package

`fmt` is a standard Go package used for:

* Printing output to the console
* Formatting strings

Example:

```go
fmt.Println("Hello World")
```

---

## 🔁 Variables & Constants

### 🔹 Variable (var)

Variables are used to store data in memory.

```go
var x int
```

* Default value is assigned automatically (`0`, `false`, etc.)
* Can be changed anytime

```go
var x int = 5
x = 10
```

Short declaration:

```go
x := 3
```

---

### 🔸 Constant (const)

Constants are immutable (cannot be changed).

```go
const x int = 5
```

❗ Must be initialized when declared

❗ Cannot use `:=` with constants

---

## 🧠 What is a Variable?

A variable is a named storage in memory.

Example logic:

* Computer stores `5` in one place
* Stores `10` in another
* Then stores result `15`

Each has:

* Name
* Value
* Type

---

## 🧪 Example (Variable vs Constant)

```go
package main

import "fmt"

func main() {
    var a int = 20
    a = 12
    fmt.Print(a) // Output: 12
}
```

```go
package main

import "fmt"

func main() {
    const a int = 20
    // a = 12 ❌ ERROR
    fmt.Print(a)
}
```

---

## 📊 Data Types in Go

Everything in computers is based on `0` and `1`, but in programming we use meaningful types.

### Basic types:

* `int`
* `float32`, `float64`
* `string`
* `bool`

Example:

```go
var a = 5
var b string = "Shahin"
var c bool = true
var d float32 = 5.32
```

---

## 🔍 Type Checking

Using `fmt`:

```go
fmt.Printf("%T", a)
```

Using `reflect`:

```go
import "reflect"

t := reflect.TypeOf(a)
fmt.Println(t)
```

---

## 🔄 Type Conversions

Go is strongly typed, so conversions must be explicit.

```go
var a int = 5
var f float64 = float64(a)
var i uint = uint(f)
```

---

## ▶️ Run the Project

```bash
go run main.go
```

---

## 🎯 Learning Progress

Topics covered:

* Go installation & setup
* package main
* fmt package
* Variables & constants
* Data types
* Type checking
* Type conversion

---

## 💡 Goal

To build a strong foundation in Go and move toward backend development.

---
🚀 Day 2 – Control Structures

📅 Date: 24.02.2026
⏱️ Study Time: 3 hours

🔀 Conditional Statements
🔹 if / else
a := 3

if a > 3 {
    fmt.Println("a is greater than 3")
} else if a == 3 {
    fmt.Println("a equals 3")
} else {
    fmt.Println("a is less than 3")
}

👉 Go-da if sadə və oxunaqlıdır
👉 () istifadə olunmur

🔹 if with short statement
if b := 30; b < 50 {
    fmt.Println("b is less than 50:", b)
}

![if-else-if](https://github.com/user-attachments/assets/3f727061-8d79-452d-84be-be4914f1a0ff)

🔁 Switch Statement

Go-da switch daha güclü və rahatdır.

🔹 Single case
a := 3

switch a {
case 1:
    fmt.Println("a is 1")
case 3:
    fmt.Println("a is 3")
default:
    fmt.Println("a is neither 1 nor 3")
}
🔹 Multiple case
score := 88

switch score {
case 56, 67:
    fmt.Println("Not 88")
case 75, 85, 88:
    fmt.Println("Approximately 88")
default:
    fmt.Println("Failed")
}
⚠️ Important Rules

❗ break yazmağa ehtiyac yoxdur (Go avtomatik edir)

❗ default – heç bir case uyğun gəlməzsə işləyir

❗ Case daxilində funksiyalar istifadə etmək olar

func test() int {
    return 10
}

a := 10

switch a {
case test():
    fmt.Println("a is 10")
}
🔄 Loops (for, range)
🔹 Basic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
🔹 While-like loop
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
🔹 Range loop
nums := []int{1, 2, 3}

for index, value := range nums {
    fmt.Println(index, value)
}
⚠️ Basic Error Handling

Go-da error handling explicit yazılır:

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

İstifadə:

result, err := divide(10, 0)

if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}

👨‍💻 Author: Shahin Quliyev
