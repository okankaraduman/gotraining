package models

import "fmt"

// Struct -> Manager

type Manager struct {
	Employee
	Projects  []string
	Locations []string
}

func (m Manager) PrintDetails() {
	m.Employee.PrintDetails()
	fmt.Println("Projects: ")
	for _, v := range m.Projects {
		fmt.Println("->", v)
	}
	fmt.Println("Managing teams for locations: ")
	for _, v := range m.Locations {
		fmt.Println(v)
	}
}
