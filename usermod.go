package main

import (
	"database/sql"
	"log"
	"net/http"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	rdb, _ := sql.Open("sqlite3", "./EPOS.db")
	defer rdb.Close()
	log.Printf("POST request (Create User) recieved", r.RemoteAddr)
	createUser(rdb, "1", username, password)
	fmt.Fprintf(w, "success")
}

func createUser(db *sql.DB, user_id string, username string, password string) {
	log.Println("Attempting creation of new user record.")
	insertUserStatement := `INSERT INTO users(user_id, username, password) VALUES (?,?,?)`
	statement, err := db.Prepare(insertUserStatement) // Prepare statement to avoid injection
	fmt.Printf("%s", insertUserStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(user_id, username, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("User %s created successfully", username)
}
