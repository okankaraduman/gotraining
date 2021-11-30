package models

import "fmt"

// Struct -> Team

type TeamMember interface {
	PrintName()
	PrintDetails()
}

type Team struct {
	Name, Description string
	TeamMembers       []TeamMember
}

func (t Team) printTeamDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members: ")
	for _, v := range t.TeamMembers {
		v.PrintName()
		v.PrintDetails()
	}
}
