package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

//streaming
func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)

	}
	defer conn.Close()
	bytes, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytes))

}
