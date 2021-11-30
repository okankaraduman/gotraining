package models

import "fmt"

//Struct -> Developer

type Developer struct {
	Employee //type embedding for composition
	Skills   []string
}

func (d Developer) PrintDetails() {
	d.Employee.PrintDetails()
	fmt.Println("Technical Skills:")
	for _, v := range d.Skills {
		fmt.Println("->", v)
	}
}
