package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	records, err := ReadData("data.csv")
	if err != nil {
		log.Fatal()
	}
	for _, record := range records {
		user := User{
			FirstName: record[0],
			LastName:  record[1],
			Job:       record[2],
		}
		fmt.Printf("%s %s is a %s \n", user.FirstName, user.LastName, user.Job)
	}
}

type User struct {
	FirstName string
	LastName  string
	Job       string
}

func ReadData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	//ilk satırı atla!
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}
	return records, nil

}
