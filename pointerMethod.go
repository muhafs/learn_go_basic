package main

import "fmt"

type Phone struct {
	Name  string
	Price int
}

//? when creating methods its recomended to use pointer ( * )
func (phone *Phone) ChangePrice(newPrice int) {
	phone.Price = newPrice
}

func main() {
	var iPhone15Pro Phone = Phone{"Iphone 15 Pro", 5000}
	fmt.Println("Phone :", iPhone15Pro.Name, "| Price :", iPhone15Pro.Price)

	iPhone15Pro.ChangePrice(4500)
	fmt.Println("Phone :", iPhone15Pro.Name, "| Price :", iPhone15Pro.Price)
}
