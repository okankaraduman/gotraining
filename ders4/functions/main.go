package main

import (
	"fmt"
	"strconv"
)

func main() {
	//printHello()

	//t := "Hello Go!"
	//printX(t)

	//t2 := "Okan!"

	//fmt.Println(login(t2))

	//name, surname := loginTwo("Okan", "Karaduman")
	//
	//fmt.Printf("Name: %s\nSurname: %s\n", name, surname)
	//x := add(2, 3)
	//w(numConverter(x))

	//message := "Hello"
	//sayHello(&message) // pointer işlemi yapacaksak fonksiyonda parametrelerde * bulunmalı.
	//fmt.Println(message)

	//result, lenOfArray := addDynamic(2, 3, 4, 5)
	//fmt.Printf("Result:%d\nLen Of Array:%d ", result, lenOfArray)

	//named results
	//fmt.Println(split(5))
	//fmt.Println(addDynamicNamedResults(1, 2, 3, 4))

	//closure bir çeşit anonim fonksiyondur
	//parameter as a function, isimlendirme yoksa anonim functiondır
	//a, b := 5, 8
	//fn := func(sum int) (int, int) {
	//	x := sum * a / b
	//	y := sum - x
	//	return x, y
	//}
	//
	//xx, yy := fn(10)
	//fmt.Println(xx, yy)
	//
	//splitValues(fn)
	//n1 := incrementCounter()
	//
	//fmt.Println("n1 increment counter #1", n1())
	//fmt.Println("n1 increment counter #1", n1())
	//n2 := incrementCounter()
	//fmt.Println("n2 increment counter #1", n2())
	//fmt.Println("n2 increment counter #1", n2())

}

func incrementCounter() func() int {
	initiliazedNumber := 0
	return func() int {
		initiliazedNumber++
		return initiliazedNumber
	}
}

func splitValues(f func(sum int) (int, int)) {
	x, y := f(40)
	fmt.Println(x, y)
}

func addDynamicNamedResults(terms ...int) (numTerms int, sum int) { //Dynamic number array with named results
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return numTerms, sum

}
func split(sum int) (x int, y int) { //geri dönüş tanımlandı zaten altta := kullanmıyoruz , named results
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func addDynamic(terms ...int) (int, int) { //Dynamic number array.
	result := 0
	for _, term := range terms {
		result += term
	}
	return result, len(terms)

}

func sayHello(msg *string) {
	fmt.Println(*msg)
	*msg = "Hi" //message'in değerini değiştir.
}

func numConverter(x int) string {
	return strconv.Itoa(x)
}

func w(p string) {
	fmt.Println(p)
}

func add(x, y int) int { ///(x int, y int) 'de olabilir. Clean code açısından hepsinin tipini ayrı yazmak gerekiyor.
	return x + y
}

//Student Examples
func printHello() {
	fmt.Println("Hello World!")
}

func printX(y string) {
	fmt.Println(y)
}

func login(uname string) string {
	return "Welcome! " + uname
}

func loginTwo(uname, surname string) (string, string) {
	return uname, surname
}
