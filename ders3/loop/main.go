package main

import (
	"fmt"
	"os"
)

func main() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
