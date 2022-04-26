package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name   string
	Price  string
	Errors map[string]string
}

// Validate form input
func (p *Product) validateProduct() bool {
	p.Errors = make(map[string]string)
	if strings.TrimSpace(p.Name) == "" {
		p.Errors["name"] = "Please enter a product name.\n"
	}

	if strings.TrimSpace(p.Price) == "" {
		p.Errors["price"] = "No Price\n"
	}

	return len(p.Errors) == 0
}

// HTTP handler for /create form inputs
func productCreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	p := Product{
		Name:  r.FormValue("name"),
		Price: r.FormValue("price"),
	}

	valid := p.validateProduct()
	if valid != true {
		log.Printf("User validation failed: %s", p.Errors)
		for key, element := range p.Errors {
			fmt.Fprintln(w, key, element)
		}
	} else {
		rdb, _ := sql.Open("sqlite3", db)
		defer rdb.Close()
		log.Printf("POST request (Create Product) recieved (%s)", r.RemoteAddr)
		createProduct(rdb, p.Name, p.Price)
	}
}

// Inserts a new Product (name, price) into the database
func createProduct(db *sql.DB, name string, price string) {
	log.Println("Attempting creation of new product record.")
	insertProductStatement := `INSERT INTO products(name, price) VALUES (?,?)`
	statement, err := db.Prepare(insertProductStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name, price)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("Product %s created successfully", name)
}
