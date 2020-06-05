package main

import "fmt"

func main() {
	// init integer array with values.
	x := [4]int{5, 10, 15, 20}
	fmt.Println(x[0])

	// init integer array without values.
	var y [2]int
	fmt.Println("set", y)
	fmt.Println("get", y[0])

	z := [2]int{5, 10}
	fmt.Println("get", z[1])

	// init 2 dimension array without values.
	var twoDimension [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimension[i][j] = i+j
		}
	}
	fmt.Println("get 2 dimesion array:", twoDimension)
}