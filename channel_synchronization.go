package main

import (
	"fmt"
	"time"
)

func worker1st(done chan bool) {
	fmt.Println("working 1st...")
	time.Sleep(time.Second * 2)
	fmt.Println("done 1st")

	done <- true
}

func worker2nd(done chan bool) {
	fmt.Println("working 2nd...")
	time.Sleep(time.Second * 4)
	fmt.Println("done 2nd")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker1st(done)
	go worker2nd(done)
	<- done
	<- done
}