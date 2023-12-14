package main

// import "fmt"

// func wordFilter(word string) string {
// 	if word == "Poop" {
// 		return "***"
// 	} else {
// 		return word
// 	}
// }

// //? type Filter func(string) string
// //? func sayFilteredHello(name string, filter Filter) { ...

// func sayFilteredHello(name string, filter func(string) string) {

// 	filteredName := filter(name)

// 	fmt.Println("Welcome", filteredName)
// }

// func main() {
// 	sayFilteredHello("Poop", wordFilter)
// 	sayFilteredHello("Salim", wordFilter)

// 	nameFilter := wordFilter
// 	sayFilteredHello("Qaswarah", nameFilter)
// 	sayFilteredHello("Poop", nameFilter)
// }
