package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Merhaba!")
	fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", handler) // cihanozhan.com/

	fmt.Println("Web server.")

	http.ListenAndServe(":8000", nil) // localhost:8000
}
