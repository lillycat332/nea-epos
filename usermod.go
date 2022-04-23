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
	Username       string
	HashedPassword string
	Privilege      string
	Errors         map[string]string
}

func (user *User) validateUser() bool {
	user.Errors = make(map[string]string)
	if strings.TrimSpace(user.Username) == "" {
		user.Errors["username"] = "Please enter a username.\n"
	}

	if strings.TrimSpace(user.HashedPassword) == "" {
		user.Errors["password"] = "No Password\n"
	}

	priv, err := strconv.Atoi(user.Privilege)
	if err != nil || priv > 2 || priv < 0 {
		user.Errors["privilege"] = "Privilege must be a numerical value between 0 and 2.\n"
	}
	return len(user.Errors) == 0
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if len(r.FormValue("password")) <= 0 {
		fmt.Fprintf(w, "Failed to create user - a password is required!")
	} else {
		pw, err := hashPassword(r.FormValue("password"))
		if err != nil {
			log.Printf("Failed to create user - hashing password failed!")
		}
		u := User{
			Username:       r.FormValue("username"),
			HashedPassword: pw,
			Privilege:      r.FormValue("privilege"),
		}

		valid := u.validateUser()
		if valid != true {
			log.Printf("User validation failed: %s", u.Errors)
			for key, element := range u.Errors {
				fmt.Fprintln(w, key, element, "\n")
			}
		} else {
			rdb, _ := sql.Open("sqlite3", db)
			defer rdb.Close()
			log.Printf("POST request (Create User) recieved (%s)", r.RemoteAddr)
			createUser(rdb, u.Username, u.HashedPassword, u.Privilege)
		}
	}
}

func createUser(db *sql.DB, username string, password string, privilege string) {
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
