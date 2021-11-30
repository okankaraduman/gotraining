package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//f, err := os.Open("scrap.csv")
	//checkError(err)
	//
	//r := csv.NewReader(f)
	//r.Read()
	//
	//records, err := r.ReadAll()
	//checkError(err)
	//for _, row := range records {
	//	printRow(row)
	//}
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var emp Employee
	var employees []Employee

	for _, each := range csvData {
		emp.FirstName = each[0]
		emp.LastName = each[1]
		emp.Age, err = strconv.Atoi(each[2])
	}

}
func printRow(row []string) {
	log.Printf("len(row) %d\n", len(row))
	for i, col := range row {
		log.Printf("[%d]: %s\n", i, col)
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
