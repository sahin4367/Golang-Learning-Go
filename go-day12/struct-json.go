package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// func main() {

// 	user := User{
// 		Name: "Shahin",
// 		Age:  21,
// 	}

// 	data, _ := json.Marshal(user)

// 	fmt.Println(string(data))
// }

func main() {
	jsonData := `{"name" : "Shahin" , "age" : 19}`
	var user User
	json.Unmarshal([]byte(jsonData), &user)

	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
