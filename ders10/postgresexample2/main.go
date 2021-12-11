package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	//preparation
	db, err := gorm.Open("postgres", "user=postgres password=okan1905 dbname=postgres2 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	dbase := db.DB()
	defer dbase.Close()

	//app
	//db.CreateTable(&User{}) - CREATING TABLE

	//db.DropTable(&User{}) //Delete old table if it exists.
	//db.CreateTable(&User{})

	//user := User{
	//	Username:  "cihanozhan",
	//	FirstName: "Cihan",
	//	LastName:  "Özhan",
	//}
	//db.Create(&user)
	//fmt.Println(user)

	//Save user list
	//for _, user := range users {
	//	db.Create(&user)
	//}

	////Get last and first data
	//u := User{}
	//db.First(&u) // Bring the first user data
	//fmt.Println(u)
	//
	//x := User{}
	//db.Last(&x) // Bring the last user data
	//fmt.Println(x)

	////Find the user
	//u := User{Username: "sjobs"}
	//db.Where(&u).First(&u) //SELECT * FROM Users WHERE Username = "sjobs";
	//fmt.Println(u)
	////Update the last name
	//u.LastName = "Özhan"
	//db.Save(&u)

	// Get the changed object
	//user := User{Username: "sjobs"}

	////Deleting users with usernamöe sjobs
	//db.Where(&User{Username: "sjobs"}).Delete(&User{})
	//fmt.Println(user)

	//db.DropTable(&User{})
	db.CreateTable(&Category{})
	fmt.Println("Islem tmm.")

}

//Go object to PostgreSQL tab

//type User struct {
//	ID       uint
//	Username string
//}
type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
	//IsActive  bool - soft delete
	//SELECT * FROM Users Where IsActive = "T"
}
type Category struct {
	//Automatically includes id,Created At, updatedAt, deletedAt objects.
	gorm.Model
}

func (u User) TableName() string { // We can create a function to change table name user-> users now user-> Userlar
	return "Userlar"
}

var users []User = []User{

	{Username: "markzuck", FirstName: "Mark", LastName: "Zuck"},
	{Username: "bgates", FirstName: "Bill", LastName: "Gates"},
	{Username: "sjobs", FirstName: "Steve", LastName: "Jobs"},
	{Username: "lellison", FirstName: "Larry", LastName: "Ellison"},
}
