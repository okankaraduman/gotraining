package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//Method ve URI
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Printf("Method -> \n", m)
			fmt.Printf("URI -> ", u)
		} else {
			//headers bitti.
			if ln == "" {
				break
			}
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `TYPE html><html lang"en"> <head> <meta charset="UTF-8"> <title> OK Website! </title>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
