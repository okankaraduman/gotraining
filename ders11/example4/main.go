package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/users/{id:[0-9]+}", get_user)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func get_user(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageConcat := "Kullanıcı ID: " + id

	message := API{messageConcat}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Hata oluştu.")
	}
	fmt.Print(w, string(output))
}

type API struct {
	Message string `json:"message"`
}
