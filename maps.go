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

	delete(m, "x")
	fmt.Println(m)
}