package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	. "github.com/valdym/golang_bootcamp/ders12/templating2/models"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	page := Page{ID: 5, Name: "User List", Description: "User List", URI: "/users"}
	users := loadUsers()
	interests := loadInterest()
	interestMappings := loadInterestMappings()

	var usersResult []User
	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		usersResult = append(usersResult, user)
	}
	viewModel := UserViewModel{Page: page, Users: usersResult}

	t, _ := template.ParseFiles("template/page.html")
	t.Execute(w, viewModel)
}

// --- LOAD Processes

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadUsers() []User {
	bytes, _ := ioutil.ReadFile("data/users.json")
	var data []User
	json.Unmarshal(bytes, &data)
	return data
}

func loadInterest() []Interest {
	bytes, _ := ioutil.ReadFile("data/interests.json")
	var data []Interest
	json.Unmarshal(bytes, &data)
	return data

}
func loadInterestMappings() []InterestMapping {
	bytes, _ := ioutil.ReadFile("data/userInterestMappings.json")
	var data []InterestMapping
	json.Unmarshal(bytes, &data)
	return data
}
