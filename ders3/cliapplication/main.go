package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

//flag

func main() {
	//name := flag.String("name", "Gopher", "name of the gopher")
	//age := flag.Int("age", 2, "age of the gopher")
	//cute := flag.Bool("cute", false, "is the gopher cute")
	//flag.Parse()
	//fmt.Printf("Gopher Stats\n Name : %s\nAge: %d\nCute: %t\n", *name, *age, *cute)

	//2

	//Usage go run main.go -name okan -age 24 -cute true

	var (
		name string
		age  int
		cute bool
	)

	flag.StringVar(&name, "name", defaultName(), "name of the gopher")
	flag.StringVar(&name, "n", defaultName(), "name of the gopher(shorthand)")
	flag.IntVar(&age, "age", defaultAge(), "age of the gopher")
	flag.IntVar(&age, "a", defaultAge(), "age of the gopher(shorthand)")
	flag.BoolVar(&cute, "cute", defaultCuteness(), "name of the gopher")
	flag.BoolVar(&cute, "c", defaultCuteness(), "name of the gopher(shorthand)")
	flag.Parse()
	fmt.Printf("Gopher Stats\n Name : %s\nAge: %d\nCute: %t\n", name, age, cute)

	//Usage : go run main.go -n okan -a 24 -c true

}

func defaultName() string {
	if os.Getenv("GOPHER_DEFAULT_NAME") != "" {
		return os.Getenv("GOPHER_DEFAULT_NAME")
	}
	return "Gopher"
}
func defaultAge() int {
	if os.Getenv("GOPHER_DEFAULT_AGE") != "" {
		age, err := strconv.Atoi(os.Getenv("GOPHER_DEFAULT_AGE"))
		if err == nil {
			return age
		}
	}
	return 7
}
func defaultCuteness() bool {
	if os.Getenv("GOPHER_DEFAULT_CUTENESS") != "" {
		cuteness, err := strconv.ParseBool(os.Getenv("CUTENESS"))
		if err == nil {
			return cuteness
		}
	}
	return true
}
