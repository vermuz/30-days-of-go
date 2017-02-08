// https://golang.org/doc/effective_go.html#maps
// https://golang.org/doc/effective_go.html#blank

// maps in go are key value based data structures similar to a hash map

package main

import "fmt"

func main() {
	// you can declare maps with the composite-literal
	x := map[string]string{"key": "value"}
	fmt.Println("x:", x)

	// keys can be of any type
	y := map[int]string{1: "value"}
	fmt.Println("y:", y)

	// we can also use make
	z := make(map[string]bool)

	// set keys
	z["key"] = true
	z["anotherKey"] = true
	fmt.Println("length of z:", len(z))

	// read or update values
	z["key"] = false
	fmt.Println("z:", z)

	// delete keys
	delete(z, "anotherKey")
	fmt.Println("length of z:", len(z))

	// just like slices we can loop over maps
	loop := map[string]int{"1": 1, "2": 2, "3": 3}
	for index := range loop {
		fmt.Println("loop:", loop[index])
	}

	// the second return value in maps is a bool
	// indicating if the map has the key or not
	val, hasKey := x["key"]
	fmt.Println("key exists:", val, hasKey)

	// the value can also be ignored using blank identifier
	_, mapHasKey := x["!key"]
	fmt.Println("key exists:", mapHasKey)
}