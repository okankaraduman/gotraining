package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func handler(w http.ResponseWriter, r *http.Request) {
	//String builder : String Birleştirme için
	var builder bytes.Buffer
	builder.WriteString("İleri Seviye T-SQL Programlama güzel bir kitaptır. Yazarını tanırım. İyi çocuktur.")
	builder.WriteString("704 Sayfa\n")
	builder.WriteString("ISBN: 8.956.487.548\n")
	builder.WriteString("Fiyat: 40.00TL\n")
	builder.WriteString("Boyut: 15x21\n")
	builder.WriteString("Baskı: 2\n")

	uri := "www.cihanozhan.com/yazilimcilar-icin-ileri-seviye-t-sql-programlama"

	page := Page{
		Title:           "Kitap: İleri Seviye T-SQL Programlama",
		Author:          "Cihan Özhan",
		Header:          "İleri Seviye T-SQL Programlama",
		PageDescription: "Bu kitabın tanıtım sayfasıdır",
		Content:         builder.String(),
		URI:             "http://" + uri,
	}
	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
