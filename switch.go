package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	switch i {
	case 3:
		fmt.Println("three")
	case 5:
		fmt.Println("five")
	case 7:
		fmt.Println("seven")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("weekdays")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before Noon")
	default:
		fmt.Println("after Noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("boolean type")
		case int:
			fmt.Println("integer type")
		default:
			fmt.Println("unknown type")
		}
	}

	whatAmI(true)
	whatAmI(5)
	whatAmI("test")
}