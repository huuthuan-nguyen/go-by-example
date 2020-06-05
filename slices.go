package main

import "fmt"

func main() {
	s := make([]string , 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(len(s))

	for i := 0; i < 3; i++ {
		fmt.Println(s[i])
	}

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s)

	w := make([]string, 10)
	copy(w, s)

	fmt.Println(w[4])

	fmt.Println(s[:3])
	fmt.Println(s[0:3])

	twoDimensionSlice := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLength := i+1;
		twoDimensionSlice[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoDimensionSlice[i][j] = i+j
		}
	}

	fmt.Println(twoDimensionSlice)
}