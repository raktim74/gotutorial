package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

// func main() {
// 	fmt.Println("starting")
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	err1 := db.Ping()

// 	if err1 != nil {
// 		log.Fatal(err1)
// 	}
// 	fmt.Println("Successfully connected")
// }

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
}

var db *sql.DB

func main() {
	var err error
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	db, err = sql.Open("postgres", "user=postgres password=root dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err1 := db.Ping()
	if err1 != nil {
		log.Fatal(err1)
	}
	http.HandleFunc("/users", getUsers)
	fmt.Println("Serever running at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from tgt.newtable")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Password, &u.Gender)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	json.NewEncoder(w).Encode(users)

}
