package main

import "fmt"

// interface
type speaker interface {
	speak() string
}

type english struct {
	name string
}

type chinese struct {
	name string
}

// interface implementation
func (english) speak() string {
	return "Hello World!"
}

// interface implementation
func (*chinese) speak() string {
	return "你好世界"
}

// polimorf function
func sayHi(spr speaker) {
	fmt.Println(spr.speak())
}

func main() {
	var speaker speaker
	var english english
	var chainese chinese

	speaker = english
	sayHi(speaker)

	speaker = &chainese
	sayHi(speaker)
}
