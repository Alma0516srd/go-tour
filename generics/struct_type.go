package main

import (
	"errors"
	"fmt"
)

type stack[T any] struct {
	Data []T
}

func (st *stack[T]) push(value T) {
	st.Data = append(st.Data, value)
}

func (st *stack[T]) pop() (T, error) {
	var zero T
	if len(st.Data) == 0 {
		return zero, errors.New("No elements in stack")
	}
	return st.Data[0], nil
}

func main() {
	var st stack[int]
	st.push(10)
	st.push(20)
	st.push(30)

	data, err := st.pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
