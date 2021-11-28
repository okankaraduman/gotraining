package main

import (
	"fmt"
	"strconv"
)

// Polymoprhism(Çok biçimlilik)
//Interface
//Type Embeddings

func main() {

	// Ferrari
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1.4
	ferr.Special = true
	// fmt.Println(ferr.Information())

	// Mercedes
	merc := new(Mercedes)
	merc.Brand = "Mercedes"
	merc.Model = "CLX"
	merc.Color = "Black"
	merc.Speed = 290
	merc.Price = 0.4
	// fmt.Println(merc.Information())

	CarExecute(ferr)
	CarExecute(merc)
}

func CarExecute(c Carface) {
	fmt.Println("\n")
	fmt.Println("Araç Bilgi: \n" + c.Information())
	fmt.Println("\n")

	msg := ""

	isRun := c.Run()
	if isRun {
		msg = "çalışıyor"
	} else {
		msg = "çalışmıyor"
	}
	fmt.Println("Araç " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "durdu"
	} else {
		msg = "durmuyor, fren tutmuyor!"
	}
	fmt.Println("Araç " + msg + ".")
}

// Struct'lar

type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

type SpecialProduction struct {
	Special bool
}

// Struct Nesneleri : Ferrari

type Ferrari struct {
	Car
	SpecialProduction
	/*
	   Type Embeddings yöntemi.
	   Object Composition(Inheritance)
	*/
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (f *Ferrari) Information() string {
	retVal := "\t" + f.Brand + " " + f.Model + "\n" + "\t" + "Color: " + f.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if f.Special {
		retVal += "\n" + "\t" + "Special: " + add
	}
	return retVal
}

// Struct Nesneleri : Lamborghini

type Lamborghini struct {
	Car
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (f *Lamborghini) Information() string {
	retVal := "\t" + f.Brand + " " + f.Model + "\n" + "\t" + "Color: " + f.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if f.Special {
		retVal += "\n" + "\t" + "Special: " + add
	}
	return retVal
}

// Struct Nesneleri : Mercedes

type Mercedes struct {
	Car
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (f *Mercedes) Information() string {
	return "\t" + f.Brand + " " + f.Model + "\n" + "\t" + "Color: " + f.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', -1, 64) + " Million"
}
