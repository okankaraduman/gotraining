package main

import (
	"fmt"
	"net/http"
)

func main() {
	var h Human
	err := http.ListenAndServe("localhost:8080", h)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// ----------------->

type Human struct {
	FName string
	LName string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.FName = "Cihan"
	hum.LName = "Özhan"
	hum.Age = 33

	r.ParseForm()
	// fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println("URL_Long", r.Form["url_long"])

	//fmt.Fprint(w, "<table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yaş</b></td></tr><tr><td>"+hum.FName+"</td><td>"+hum.LName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr><tr></tr><tr></tr><tr><td><b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")
}
