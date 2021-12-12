package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {

	// cihanozhan.com/api
	apiRoot := "/api"

	// .../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		// message := "API Home"
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{5, "Ali", "Veli", 27}
		// message := user
		output, err := json.Marshal(user)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, FirstName: "Zeynep", LastName: "Kesici", Age: 32},
			{ID: 2, FirstName: "Murtaza", LastName: "Biçici", Age: 56},
			{ID: 5, FirstName: "Adem", LastName: "Hazreti", Age: 74},
		}
		// message := users
		output, err := json.Marshal(users)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

type API struct {
	Message string `json:"message"`
}
