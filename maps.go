package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["x"] = 1
	m["y"] = 2
	m["z"] = 3
	
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println(len(m))

	delete(m, "z")
	fmt.Println(m)

	// underscore to indicate if the key in map does exists or not.
	_, t := m["z"]
	fmt.Println(t)

	n := map[string]int{"a": 1, "b": 2}
	fmt.Println(n)
}