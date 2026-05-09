package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	// XMLName xml.Name `xml:"user"` --> // Bu root tag yaradır:
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {

	user := User{
		Name: "Vusal",
		Age:  33,
	}

	data, err := xml.Marshal(user)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

//XML = eXtensible Markup Language
// <user>
//     <name>Vusal</name>
//     <age>33</age>
// </user>

// JSONA oxsarligi var
// {
//   "name": "Vusal",
//   "age": 33
// }
