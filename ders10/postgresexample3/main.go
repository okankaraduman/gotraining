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

	// app

	// db.DropTable(&User{}) // Delete old table if it exists.
	// db.CreateTable(&Category{})
	db.CreateTable(&User{})

	fmt.Println("Islem tmm.")
}

type Category struct {
	// Automatically includes id, createdAt, updatedAt, deletedAt objects.
	gorm.Model
}

// INSERT INTO Users(id, user_id, username, "FirstName", "LastName")
// VALUES(3, 44, 'murti', 'Murtaza', 'Abuzittin')
type User struct {
	ID int // GORM doesn't like 'id'. Should be 'ID'.
	// Creates a primary key for UserID.
	UserID int `gorm:"primary_key"` //burda yatık tırnak kullanılmalı
	// Creates constrains for Username
	// -> 15 character max limit and not be passed a the blank
	Username string `sql:"type:VARCHAR(15);not null"`
	// Creates constraints for FirstName
	// -> 100 character max limit, Not be passed a the blank, Column name will not be FName, will be FirstName
	FName string `sql:"size:100;not null" gorm:"column:FirstName"`
	// Creates consstraints for LastName
	// -> Unique index/constraint, Not be passed a the blank, Default value is 'Unknown', Column name will not be LName, will be LastName
	LName     string `sql:"unique;unique_index;not null;DEFAULT:'Unknown'" gorm:"column:LastName"`
	Count     int    `gorm:"AUTO_INCREMENT"`
	TempField bool   `sql:"-" ` // Ignore a Field
}
