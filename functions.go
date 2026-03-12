package main

import "fmt"

type User struct {
	name string
	age  int
}

func createUser(name string, age int) (*User, error) {
	return &User{name, age}, nil
}

func main() {
	newUser, err := createUser("Bob", 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newUser)

	_, err = createUser("test", 43334)
	if err != nil {
		fmt.Println(err)
	}
}
