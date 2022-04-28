package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// loginHandler handles the /login form input and writes the response to the client (Login failed/successful)
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	log.Printf("Login request recieved from %s", r.RemoteAddr)
	userName := r.FormValue("username")
	unhashedPassword, err := hashPassword(r.FormValue("password"))
	if err != nil {
		log.Println(err.Error())
	}
	if testPass(userName, unhashedPassword) {
		w.WriteHeader(http.StatusOK)
		enableCors(&w)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Login Successful")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		enableCors(&w)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Login Failed")
	}
}

// TestPass tests the password against the database and returns true if the password matches
func testPass(userName string, unhashedPassword string) bool {
	rdb, _ := sql.Open("sqlite3", db)
	defer rdb.Close()
	u := authQueryUser(rdb, userName)
	passIsValid := false
	if checkPasswordMatch(u, unhashedPassword) {
		passIsValid = true
	}
	return passIsValid
}

// authQueryUser performs an SQL query to get the user's hashed password from the database, returning it if successful
func authQueryUser(db *sql.DB, userName string) string {
	var p string
	rows, err := db.Query("SELECT password FROM users WHERE username = ?", userName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var password string
		err := rows.Scan(&password)
		if err != nil {
			log.Fatalln(err.Error())
		}
		p = password
	}
	return p
}
