package main

import "fmt"

const SERVER_IP = "178.56.543.11"
const SERVER_PORT int32 = 8090

func main() {
	fmt.Println("Server IP: ", SERVER_IP)
	fmt.Println("Server Port: ", SERVER_PORT)

	minutesInValue := 7777 / 60
	fmt.Println("Minute Soin Value: ", minutesInValue) //выведет с округлленеиме до ближ целого

}
