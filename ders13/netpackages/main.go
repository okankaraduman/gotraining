package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

//streaming
func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "\nHello From TCP Server!\n")
		fmt.Fprint(conn, "Go nasıl ama?\n")
		fmt.Fprintf(conn, "%v", "Güzel gidiyyaa!")

		conn.Close()
	}
}
