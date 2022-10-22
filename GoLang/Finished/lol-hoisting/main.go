package main

import "fmt"

var partyTime = "Artie"
var lameTime = "Tyler"

func main() {
	weHoistingOutHere(&partyTime)
}

func weHoistingOutHere(text *string) {
	fmt.Println(*text)

	text = &lameTime

	if 1 == 1 {
		text := "Yo like, what?"
		fmt.Println(text)
	}

	fmt.Println(*text)
}
