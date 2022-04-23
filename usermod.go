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
	user.Errors = make(map[string]string)
	if strings.TrimSpace(user.Username) == "" {
		user.Errors["username"] = "Please enter a username"
	}
	if strings.TrimSpace(user.Password) == "" {
		user.Errors["username"] = "Please enter a password"
	}
	return len(user.Errors) == 0
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	priv, err := strconv.Atoi(r.FormValue("privilege"))
	if err != nil {
		fmt.Fprintf(w, "Atoi err: %v. Enter a valid integer.", err)
	}
	u := User{
		Username:  r.FormValue("username"),
		Password:  r.FormValue("password"),
		Privilege: priv,
	}
	u.validateUser()
	rdb, _ := sql.Open("sqlite3", "./EPOS.db")
	defer rdb.Close()
	log.Printf("POST request (Create User) recieved (%s)", r.RemoteAddr)
	createUser(rdb, u.Username, u.Password, u.Privilege)
	fmt.Fprintf(w, "success")
}

func createUser(db *sql.DB, username string, password string, privilege int) {
	log.Println("Attempting creation of new user record.")
	insertUserStatement := `INSERT INTO users(username, password, privilege) VALUES (?,?,?)`
	statement, err := db.Prepare(insertUserStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(username, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("User %s created successfully", username)
}
