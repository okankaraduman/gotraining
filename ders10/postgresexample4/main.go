package main

// go get -u github.com/jinzhu/gorm
// go get -u github.com/lib/pq

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {

	// preperation
	db, err := gorm.Open("postgres", "user=postgres password=okan1905 dbname=postgres2 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//db.DropTable(&User{}) //Delete old table if it exists.
	//db.CreateTable(&User{})
	//// Reflection ile yapıldı(Fields)
	//for _, f := range db.NewScope(&User{}).Fields() {
	//	fmt.Println(f.Name)
	//}

	//Index Creation
	//db.Model(&User{}).AddIndex("idx_first_name", "first_name")
	//db.Model(&User{}).AddUniqueIndex("idx_email", "email")

	// Index Deletion
	//db.Model(&User{}).RemoveIndex("idx_email")

	db.DropTable(&Calendar{}) //Delete old table if it exists.
	db.CreateTable(&Calendar{})
	db.DropTable(&User{}) //Delete old table if it exists.
	db.CreateTable(&User{})

	db.Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
		},
	})

	u := User{}
	c := Calendar{}
	db.First(&u).Related(&c, "calendar")

	fmt.Println(u)
	fmt.Println(c)

	fmt.Println("Islem tmm")
}

type User struct {
	//Model gorm.Model yapınca field olarak Model dondu onun yerine aşağıdaki gibi tanımladık
	gorm.Model
	Username   string
	FirstName  string
	LastName   string
	Calendar   Calendar
	CalendarID uint
}

type Calendar struct {
	gorm.Model
	Name string
}
