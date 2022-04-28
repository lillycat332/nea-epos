package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name   string
	Price  string
	Errors map[string]string
}

func productReader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	rdb, _ := sql.Open("sqlite3", db)
	defer rdb.Close()
	log.Printf("POST request (Get Products) recieved (%s)", r.RemoteAddr)
	json.NewEncoder(w).Encode(queryProducts(rdb))
}

// Validate form input
func (p *Product) validateProduct() bool {
	p.Errors = make(map[string]string)
	if strings.TrimSpace(p.Name) == "" {
		p.Errors["name"] = "Please enter a product name.\n"
	}

	if strings.TrimSpace(p.Price) == "" {
		p.Errors["price"] = "Please Enter a price.\n"
	}

	return len(p.Errors) == 0
}

// HTTP handler for /createProduct form inputs
func productCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func queryProducts(db *sql.DB) []Product {
	var products []Product
	rows, err := db.Query("SELECT productName, price FROM products")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var productName string
		var price int
		err := rows.Scan(&productName, &price)
		if err != nil {
			log.Fatalln(err.Error())
		}
		pr := strconv.Itoa(price)
		p := Product{
			Name:  productName,
			Price: pr,
		}
		products = append(products, p)
	}
	return products
}
