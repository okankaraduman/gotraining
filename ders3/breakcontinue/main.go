package main

import (
	"fmt"
	"time"
)

func main() {
	// break

	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i is:", i)
		i++
	}

	s := 0
	for {
		if s >= 5 {
			break
		}
		fmt.Println("s is:", s)
		s++
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("Break happened")

	// continue

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
