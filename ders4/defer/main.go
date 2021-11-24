package main

import "fmt"

//defer

// kendisini çevreleyen fonksiyon dönene kadar fonksiyonun çalışmasını erteler.

var isConnected bool = false

func main() {
	fmt.Printf("Connection Open: %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connection Open: %v\n", isConnected)

}

func databaseProcessing() {
	connect()
	fmt.Println("Deferring disconnect!")
	defer disconnect() // wait until databaseProcessing() function finishes
	defer disconnect2()
	// .....
	// .....
	// .....
	// .....
	fmt.Println("Do something...")
	// .....
	// .....
	// .....
	// .....
	fmt.Println("Connection Open: End of scope")
}
func connect() {
	isConnected = true
	fmt.Println("Connected to Database!")
}
func disconnect() {
	isConnected = false
	fmt.Println("Disconnected from Database!")
}
func disconnect2() {
	isConnected = false
	fmt.Println("Disconnected from Database! 2")
}
