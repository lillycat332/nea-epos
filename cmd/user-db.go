package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// User is a struct containing the Username (string)
// HashedPassword (string), Privilege (string),
// and any errors (map[string]string) of a given user.
// These are used to represent users before being entered
// into the database.
type User struct {
	Username       string
	HashedPassword string
	Privilege      string
	Errors         map[string]string
}

// validateUser validates a form input for a User.
// Some validation has to take place outside of this
// method, because of the hashing.
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

// HTTP handler for /readUsers requests
func userReader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	EnableCors(&w)
	w.Header().Set("Content-Type", "text/plain")
	rdb, _ := sql.Open("sqlite3", db)
	defer rdb.Close()
	log.Printf("POST request (Get Users) recieved (%s)", r.RemoteAddr)
	io.WriteString(w, strings.Join(queryUsers(rdb), ", "))
}

// HTTP handler for /create form inputs
func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	log.Printf("POST request (Create User) recieved (%s)", r.RemoteAddr)

	if len(r.FormValue("password")) <= 0 {
		fmt.Fprintf(w, "Failed to create user - a password is required!")
	} else {
		pw, err := HashPassword(r.FormValue("password"))
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
				fmt.Fprintln(w, key, element)
			}
		} else {
			rdb, _ := sql.Open("sqlite3", db)
			defer rdb.Close()
			createUser(rdb, u.Username, u.HashedPassword, u.Privilege)
		}
	}
}

// createUser creates a user by inserting
// into the database (*sql.DB).
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
	arr := queryUsers(db)
	log.Printf("INFO: Users in DB are: %s", arr)
}

// queryUsers queries the database for users matching
// a simple SQL query.
func queryUsers(db *sql.DB) []string {
	var users []string
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatalln(err.Error())
		}
		users = append(users, username)
	}
	return users
}
