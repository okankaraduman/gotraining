package main

/*
CREATE TABLE cities(
  id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    population INT
);
*/

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valdym/ders8/mysqlproject/models"
)

var db *sql.DB

func main() {

	db := conn()

	// check_version()

	// select_all_rows(db)

	// insert_row(db)

	// select_by_id(db, 11)

	delete_rows(db, 3)
}

func conn() *sql.DB {
	db, err := sql.Open("mysql", "root:1234567@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func check_version() {
	// username:pwd@tcp(server:port)/dbname
	db, err := sql.Open("mysql", "root:1234567@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(sql.Drivers())

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(version)
}

func select_all_rows(db *sql.DB) {
	response, err := db.Query("SELECT * FROM cities")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	for response.Next() {
		var city models.City
		err := response.Scan(&city.Id, &city.Name, &city.Population)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", city)
	}
}

func insert_row(db *sql.DB) {
	sql := "INSERT INTO cities(name, population) VALUES('Istanbul',180000000)"
	result, err := db.Exec(sql)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	affected, err := result.RowsAffected()
	fmt.Printf("%d kayıt etkilendi.", affected)
}

func select_by_id(db *sql.DB, id int) {
	sql := "SELECT * FROM cities WHERE id = " + strconv.Itoa(id)
	response, err := db.Query(sql)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	if response.Next() {
		var city models.City
		err := response.Scan(&city.Id, &city.Name, &city.Population)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", city)
	} else {
		fmt.Println("No data found.")
	}
}

func delete_rows(db *sql.DB, id int) {
	sql := "DELETE FROM cities WHERE id = " + strconv.Itoa(id)
	result, err := db.Exec(sql)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d kayıt silindi.", affectedRows)
}

// DELETE FROM X WHERE id IN(1,2,3); INSERT ...
// MARS = Mupliple Active Result Set

// SELECT .. FROM x; DROP TABLE users;
// ; DROP TABLE users;

/*

  /users/
    - GET /users/:id
      adada
    - POST
    - DELETE /users/:id
      POST ile de veri silme işlemi yapılabilir. Fakat...
  /products/
    - adada

*/

//Swagger
