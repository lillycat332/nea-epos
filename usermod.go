package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username  string
	Password  string
	Privilege int
	Errors    map[string]string
}

func (user *User) validateUser() bool {
	Errors := make(map[string]string)
	if strings.TrimSpace(user.Username) == "" {
		Errors["username"] = "Please enter a username"
	}
	if strings.TrimSpace(user.Password) == "" {
		Errors["password"] = "Please enter a password"
	}
	return len(Errors) == 0
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	priv, err := strconv.Atoi(r.FormValue("privilege"))
	if err != nil || priv > 2 || priv < 0 {
		fmt.Fprintf(w, "Privilege must be a numerical value between 0 and 2. ")
	}

	pw, err := hashPassword(r.FormValue("password"))
	if err != nil {
		log.Printf("Failed to create user - hashing password failed!")
	}

	u := User{
		Username:  r.FormValue("username"),
		Password:  pw,
		Privilege: priv,
	}

	valid := u.validateUser()
	if valid != true {
		log.Printf("User validation failed: %s", u.Errors)
		fmt.Fprintf(w, "Invalid user entry.")

	} else {
		rdb, _ := sql.Open("sqlite3", "./EPOS.db")
		defer rdb.Close()
		log.Printf("POST request (Create User) recieved (%s)", r.RemoteAddr)
		createUser(rdb, u.Username, u.Password, u.Privilege)
	}
}

func createUser(db *sql.DB, username string, password string, privilege int) {
	log.Println("Attempting creation of new user record.")
	insertUserStatement := `INSERT INTO users(username, password, privilege) VALUES (?,?,?)`
	statement, err := db.Prepare(insertUserStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(username, password, privilege)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("User %s created successfully", username)
}
