package main

import "fmt"

func main() {
	nums := []int{5, 10, 15}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	kvs := map[string]string{"a": "A", "b": "B", "c": "C"}
	for k, v := range kvs {
		fmt.Println("k => v", k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}

	for _, v := range kvs {
		fmt.Println(v)
	}

	for i, c := range "abcdefgh" {
		fmt.Println("index with character:", i, c)
	}
}