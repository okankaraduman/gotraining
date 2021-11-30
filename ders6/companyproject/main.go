package main

import (
	"time"

	. "yemeksepeti.com/ecommerce/dev-team/models"
)

//Interface, Composition, Method Overriding

func main() {
	//Gamze, Batuhan, Ayberk, Busra
	gamze := Developer{
		Employee: Employee{
			FirstName: "Gamze",
			LastName:  "Karadayı",
			Dob:       time.Date(2000, time.February, 15, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Senior Go Developer",
			Location:  "İzmir",
		},
		Skills: []string{"Go", "Docker", "Kubernetes", "C#"},
	}
	ellerhavada_batuhan := Developer{
		Employee: Employee{
			FirstName: "Batuhan",
			LastName:  "D.",
			Dob:       time.Date(1999, time.February, 18, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Junior Go Developer",
			Location:  "İzmir",
		},
		Skills: []string{"Go", "MySQL", "PHP", "ANV"},
	}
	ayberk := Developer{
		Employee: Employee{
			FirstName: "Ay",
			LastName:  "Bi",
			Dob:       time.Date(1998, time.July, 14, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Go Rapper",
			Location:  "Istanbul",
		},
		Skills: []string{"Go", "Python", "PostgreSQL", "Rust"},
	}
	mslogic := Manager{
		Employee: Employee{
			FirstName: "Bayan",
			LastName:  "Mantik",
			Dob:       time.Date(1987, time.March, 13, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Team Manager",
			Location:  "Yozgat",
		},
		Projects:  []string{"Face Detection", "e-Commerce"},
		Locations: []string{"San Fransisco", "Yozgat"},
	}
	team := Team{
		"Go",
		"Golang Engineering Team",
		[]TeamMember{gamze, ayberk, ellerhavada_batuhan, mslogic},
	}
	team.printTeamDetails()

}
