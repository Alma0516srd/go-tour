package main

import "fmt"

type keymap[T any] map[string]T

func (genMap keymap[T]) set(key string, val T) {
	genMap[key] = val
}

func (genMap keymap[T]) get(key string) (T, bool) {
	var notFound T
	val, exists := genMap[key]
	if !exists {
		return notFound, false
	}

	return val, true
}

func main() {
	genMap := make(keymap[int])
	genMap.set("test1", 1)
	genMap.set("test2", 2)
	genMap.set("test3", 3)
	get, _ := genMap.get("test1")
	getVal := get
	fmt.Println(getVal)

}
