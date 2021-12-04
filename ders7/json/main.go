package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	person := Person{
		ID:        7,
		FirstName: "Cihan",
		LastName:  "Özhan",
		UserName:  "cihanozhan",
		Gender:    "Erkek",
		Name: Name{
			Family:   "xyz",
			Personal: "Cihan",
		},
		Email: []Email{
			{ID: 1,
				Kind:    "work",
				Address: "cihan@deeplab.co"},
			{ID: 2,
				Kind:    "home",
				Address: "cihan.ozhan@hotmail.com"},
		},
		Interest: []Interest{
			{
				ID:   1,
				Kind: "C#",
			},
			{
				ID:   2,
				Kind: "Go",
			},
			{
				ID:   3,
				Kind: "Python",
			},
			{
				ID:   4,
				Kind: "SQL Server",
			},
			{
				ID:   5,
				Kind: "PostgreSQL",
			},
		},
	}
	WriteMessage("Reading Operation Started.")

	WriteMessage("Personal Fullname")
	WriteStarLine()
	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personal Email Object with Index")
	WriteStarLine()

	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarLine()

	WriteMessage("Reading Operation Ended.")
	WriteMessage("\n")
	WriteMessage("Writing Operation Started.")
	SaveJSON("person.json", person)
	WriteMessage("Writing operation end")
}

//Struct Objects
type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Kind string
}

//Function ve Method'lar

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, ind int) string {
	return p.Email[ind].Address
}
func GetPersonEmail(p *Person, ind int) Email {
	return p.Email[ind]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("*************")
}
func SaveJSON(filename string, key interface{}) {
	outFile, err := os.Create("person.json")
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error :", err.Error())
		os.Exit(1)
	}
}
