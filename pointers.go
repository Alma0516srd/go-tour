package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeAge(personStruct *Person, newAge int) {
	personStruct.age = newAge
}

func main() {

	var coins int

	coins = 20

	fmt.Println(&coins)
	fmt.Println(coins)

	coinsPointer := &coins
	fmt.Println("Address:", &coinsPointer) //адрес указателя
	fmt.Println("Value:", coinsPointer)    //Значение указателя
	fmt.Println("Pointer", *coinsPointer)  // значение на которое указывает указатель

	bob := Person{"Bob", 20}

	fmt.Println("Bob age before change ", bob.age)
	changeAge(&bob, 55) //идет передача адреса на структуру - значение менять можно

	fmt.Println("Bob age after change ", bob.age)

}
