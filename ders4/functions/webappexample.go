package main

//Anonymous example
import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(writeResponse)
	http.Handle("/", addHeader(newStatusCode(handler))) //abc.com//abc.com/
	http.Handle("/onlyX", addHeader(handler))           //abc.com/onlyx
	http.Handle("/onlyY", newStatusCode(handler))       //abc.com/onlyy
	http.Handle("/admin", adminCheck(handler))          //abc.com/admin
	http.ListenAndServe(":1234", nil)
}

func adminCheck(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("user") != "admin" {
			http.Error(w, "Not Authorized", 401)
			return
		}
		fmt.Fprintln(w, "Admin Page")
		h.ServeHTTP(w, r)
	})
}

func newStatusCode(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		h.ServeHTTP(w, r) //w,r kısaltmaları writer, request çok kullanılır
	})
}

func addHeader(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Foo", "Bar")
		h.ServeHTTP(w, r) //w,r kısaltmaları writer, request çok kullanılır
	})
}

func writeResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go!")
}
