package main

import (
	"fmt"
	"strconv"

	"github.com/valdym/userpayment/models"
	"github.com/valdym/userpayment/utils"
	//. "github.com/valdym/userpayment/models" //.'yla tanımlandığında direk fonksiyonları kullanabiliyoruz.
)

//Struct

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//// v1
	//
	//u1 := &User{
	//	ID:        5,
	//	FirstName: "Ali",
	//	LastName:  "Veli",
	//	UserName:  "aliveli",
	//	Age:       29,
	//	Pay: &Payment{
	//		Salary: 3.555,
	//		Bonus:  522,
	//	},
	//}
	//
	//fmt.Println(u1.GetFullName())
	//fmt.Println(u1.GetPayment())

	// v2
	// User Creation
	user := models.NewUser()
	user.FirstName = "Cihan"
	user.LastName = "Özhan"
	user.Age = 33
	user.UserName = "cihanozhan"
	// Payment Creation v1
	// user.Pay.Salary = 50
	// user.Pay.Bonus = 5
	// // Printing
	// fmt.Println(user.Pay)
	// fmt.Println(user.GetPayment())
	// Payment Creation v2
	user.Pay = &models.Payment{Salary: 45, Bonus: 3}
	fmt.Println(user.GetFullName())
	// fmt.Println(user.Pay)
	fmt.Println(user.GetPayment())

	//random

	randX := utils.Random(10, 99)
	fmt.Println("Random value:", strconv.Itoa(randX))

	/// ----------------------------------------------------//
	//v1

	//humX := Human{FirstName: "Murtaza", LastName: "Soylu"}
	//fmt.Println("Data : " + humX.FirstName)
	//fmt.Println("Data : " + humX.LastName)

	//humX := Human{}
	//humX.FirstName = "Veli"
	//fmt.Println("Data : " + humX.FirstName)

	//v2

	//humY := new(Human)
	//humY.FirstName = "Ayşecan"
	//humY.LastName = "X"
	//
	//fmt.Println("Data : " + humY.FirstName)

	//xx := NewHuman("CO", "Ozhan")
	//fmt.Println(xx.FirstName + " " + xx.LastName)

}

func NewHuman() *Human {
	return new(Human)
}

////Constructor function
//func NewHuman(FirstName, lastName string) *Human {
//	x := new(Human)
//	x.FirstName = FirstName
//	x.LastName = lastName
//	return x
//}
