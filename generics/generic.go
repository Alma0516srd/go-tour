package main

import (
	"encoding/json"
	"fmt"
)

type MarshalUser struct {
	Name string
	Age  int
}

func marshal[T any](val T) (string, error) {
	bytes, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func main() {
	u := MarshalUser{"Alice", 43}
	s, err := marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
