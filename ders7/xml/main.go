package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//XML to GO

func main() {
	xmlFile, err := os.Open("Employees.xml")
	if err != nil {
		fmt.Println("Opening file error :", err)
		return
	}
	defer xmlFile.Close()
	xmlData, _ := ioutil.ReadAll(xmlFile)

	//Unmarshalling
	var c Company
	xml.Unmarshal(xmlData, &c) //XML to GO Object işlemini yapan bu. Bu işlem için struct'la xml formatı birebir eşleşmek zorunda.
	//fmt.Println(c.Persons)

	// - - - - - - - - - - - - - - - - - - - >

	// 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
	//jsonData, err := json.MarshalIndent(c, "", " ")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(jsonData))

	//XML to JSON

	var person jsonPerson
	var persons []jsonPerson
	for _, value := range c.Persons {
		person.ID = value.ID
		person.FirstName = value.FirstName
		person.LastName = value.LastName
		person.UserName = value.UserName

		persons = append(persons, person)
	}

	jsonData, err := json.MarshalIndent(persons, "", " ") //Marshal also works, indent creates indent between members

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print JSON Data
	fmt.Println(string(jsonData))

	//Write to JSON file

	jsonFile, err := os.Create("Employees.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}

type jsonPerson struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

//models kısmını oluştur bunlar için!.
type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username" `
	//UserName  string   `xml:"username" json:"uname"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

func (p Person) String() string {
	return fmt.Sprintf("\t ID:%d - FirstName : %s %s -> UserName : %s \n",
		p.ID, p.FirstName, p.LastName, p.UserName)
}
