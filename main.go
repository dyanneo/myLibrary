package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// _ "github.com/go-sql-driver/mysql"  <- import when ready to use db

// Book is the main object of this app
type Book struct {
	ID         int
	Title      string
	Author     string
	FormatType string
	Location   string
}

// <- uncomment when ready to use db
func dbConn() (db *sql.DB) {
	//fmt.Println("top of dbConn")
	dbDriver := "mysql"
	dbUser := "user"
	dbPass := "pass"
	dbName := "golibrary"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Printf("conn fail: %s\n", err)
		panic(err.Error())
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("ping fail: %s\n", err)
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

// Index is the main page
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("top of Index")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM book ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var id int
		var title, author, formattype, location string
		err = selDB.Scan(&id, &title, &author, &formattype, &location)
		if err != nil {
			panic(err.Error())
		}
		book.ID = id
		book.Title = title
		book.Author = author
		book.FormatType = formattype
		book.Location = location
		res = append(res, book)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()

}

// Add a new book to the database
func Add(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("top of Add")
	tmpl.ExecuteTemplate(w, "Add", nil)
}

// Insert performs the db write
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("top of Insert")
	db := dbConn()
	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		formattype := r.FormValue("formattype")
		location := r.FormValue("location")
		insForm, err := db.Prepare("INSERT INTO book(title, author, formattype, location) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(title, author, formattype, location)
		log.Println("INSERT: Title: " + title + " | Author: " + author + " | Format: " + formattype + " | Location: " + location)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {

	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	// http.HandleFunc("/show", Show)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/print", Print)
	// http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)

}
