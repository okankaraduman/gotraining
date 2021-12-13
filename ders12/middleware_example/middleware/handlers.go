package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/valdym/gotraining/ders12/middleware_example/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConn() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully conntected!")
	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode: %v", err)
	}
	insertedID := insertUser(user)
	res := response{
		ID:      insertedID,
		Message: "User created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

/// ----->

func insertUser(user models.User) int64 {
	db := createConn()
	defer db.Close()

	sql := "INSERT INTO users (name,location,age) VALUES($1,$2,$3) RETURNING id"
	var id int64
	err := db.QueryRow(sql, user.Name, user.Location, user.Age)
	if err != nil {
		log.Fatal("Unable to execute the query : %v", err)
	}
	fmt.Printf("Inserted a single record : %v", id)
}
