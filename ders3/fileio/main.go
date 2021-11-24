package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter text: ")
	str, _ := reader.ReadString('\n') //Read until newline

	fmt.Println(str)

	fmt.Println("enter number : ")
	str, _ = reader.ReadString('\n') //Read until newline

	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("value of number : ", f)
	}
}
