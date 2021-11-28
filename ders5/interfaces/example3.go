package main

import (
  "fmt"
  "time"
)

func main() {
  // var p Person
  // p.FirstName = "Bla bla"

  p := Person{
    FirstName: "Cihan",
    LastName: "Özhan",
    Dob: time.Date(2000, time.February, 20, 0, 0, 0, 0, time.UTC),
    Email: "cihan@deeplab.co",
    Location: "Istanbul",
  }
  fmt.Println(p.PrintName())
  p.PrintDetails()
}

type Person struct {
  FirstName, LastName string
  Dob time.Time
  Email, Location string
}

func (p *Person) PrintName() string {
  return p.FirstName + " " + p.LastName
}

func (p *Person) PrintDetails() {
  fmt.Printf("Date of Birth: %s, Email: %s, Location: %s\n", p.Dob.String(), p.Email, p.Location)
}
11:09

