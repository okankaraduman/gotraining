package main

//Konfigürasyon Dosyaları

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

/*
	Konfigürasyon Yöntemleri
		- Ortam Değişkenleri
		- Object
		Konfigürasyon Dosyaları
			.txt
			.env
			.ini
			.cfg
			.json
			.xml
			.yaml
			.toml
			.boml

		DB_CONNECTION=mysql
		DB_HOST=localhost
		DB_PORT=3306
		DB_DATABASE=abc
		DB_USERNAME=def
		DB_PASSWORD=lsjvosjdvıod
*/

func main() {
	x:= Person{"Ali","Veli",55},
	y, err := yaml.Marshal(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(y))
}

type Person struct {
	FirstName string `json:"first_name" yaml:"first_name"`
	LastName  string `json:"last_name" yaml:"last_name"`
	Age       int    `json:"age" yaml:"age"`
}
