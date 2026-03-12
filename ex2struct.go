package main

import "fmt"

type user struct {
	name  string
	age   int
	email string
}

func main() {

	user := user{
		name:  "Jack",
		age:   42,
		email: "usermail@example.com",
	}
	fmt.Println("name:", user.name)
	fmt.Println("age:", user.age)
	fmt.Println("email:", user.email)

	building := struct {
		name      string
		totalArea int
		address   string
	}{
		name:      "AxisTower",
		totalArea: 420998,
		address:   "San Francisco, CA 889900",
	}
	fmt.Println("building:", building)
	fmt.Println("totalArea:", building.totalArea)
	fmt.Println("address:", building.address)
	fmt.Println("name:", building.name)

}
