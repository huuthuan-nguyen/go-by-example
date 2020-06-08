package main

import "fmt"

func plus(x int, y int) int {
	return x + y
}

func triplePlus(x, y, z int) int {
	return x + y + z
}

func main() {
	sum := plus(5, 10)
	fmt.Println(sum)

	tripleSum := triplePlus(5, 10, 15)
	fmt.Println(tripleSum)
}
