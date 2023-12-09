package main

import "fmt"

func main() {
	type CardID string

	var UwaisID CardID = "34994050023001203404"
	fmt.Println("Length ID = ", len(UwaisID))
	fmt.Println("Uwais ID = ", UwaisID)
}
