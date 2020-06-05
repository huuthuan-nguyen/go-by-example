package main

import "fmt"

func main() {
	var a int = 4

	if a%2 == 0 {
		fmt.Println("Wow!");
	} else {
		fmt.Println("Fuck");
	}

	if num := 9; num < 0 {
		fmt.Println("Less than zero")
	} else if num == 9 {
		fmt.Println("Equal nine")
	} else {
		fmt.Println("none");
	}
}