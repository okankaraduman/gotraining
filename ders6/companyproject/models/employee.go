package models

import (
	"fmt"
	"time"
)

type Employee struct {
	FirstName, LastName string
	Dob                 time.Time
	JobTitle, Location  string
}

func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date Of Birth:%s, Job:%s,Location:%s\n", e.Dob.String(), e.JobTitle, e.Location)
}
