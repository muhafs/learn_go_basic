package main

// import "fmt"

// type BlackList func(string) bool

// func registerUser(name string, blackList BlackList) {
// 	if blackList(name) {
// 		fmt.Println("You are Blocked", name)
// 	} else {
// 		fmt.Println("Welcome", name)
// 	}
// }

// func main() {
// 	blacklistAdmin := func(name string) bool {
// 		return name == "admin"
// 	}

// 	registerUser("admin", blacklistAdmin)

// 	registerUser("anjing", func(s string) bool {
// 		return s == "anjing"
// 	})

// 	registerUser("Uwais", blacklistAdmin)

// 	registerUser("Yara", func(s string) bool {
// 		return s == "anjing"
// 	})
// }
