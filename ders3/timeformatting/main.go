package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2021, time.November, 10, 20, 0, 0, 0, time.UTC)

	fmt.Printf("Go Launch at %s\n", t)
	now := time.Now()
	fmt.Printf("Go Launch at %s\n", now)

	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())
	fmt.Println("---------------")

	//Tarihe 1 gün ekleme

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	shortFormat := "06/3/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

}
