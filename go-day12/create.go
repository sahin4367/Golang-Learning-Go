package main

import "os"

func main() {
	file, err := os.Create("note.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	//Burda meqsed memoriliyin qarsinin alir ,
	//file look problemimim hell edir ve s.

	file.WriteString("go run Shahin")

}
