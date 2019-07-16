package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// Book is the main object of this app
type Book struct {
	ID         int
	Title      string
	Author     string
	FormatType string
	Location   string
	ISBN       string
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
	selDB, err := db.Query("SELECT * FROM book ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var id int
		var title, author, formattype, location, isbnnumber string
		var isbn sql.NullString
		err = selDB.Scan(&id, &title, &author, &formattype, &location, &isbn)
		if err != nil {
			panic(err.Error())
		}
		if isbn.Valid {
			isbnnumber = isbn.String
		}
		book.ID = id
		book.Title = title
		book.Author = author
		book.FormatType = formattype
		book.Location = location
		book.ISBN = isbnnumber
		res = append(res, book)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()

}

// Show a book's full details
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM book WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	for selDB.Next() {
		var id int
		var title, author, formattype, location, isbnnumber string
		var isbn sql.NullString
		err = selDB.Scan(&id, &title, &author, &formattype, &location, &isbn)
		if err != nil {
			panic(err.Error())
		}
		if isbn.Valid {
			isbnnumber = isbn.String
		}
		book.ID = id
		book.Title = title
		book.Author = author
		book.FormatType = formattype
		book.Location = location
		book.ISBN = isbnnumber
	}
	tmpl.ExecuteTemplate(w, "Show", book)
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
		isbn := r.FormValue("isbn")
		insForm, err := db.Prepare("INSERT INTO book(title, author, formattype, location, isbn) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(title, author, formattype, location, isbn)
		log.Println("INSERT: Title: " + title + " | Author: " + author + " | Format: " + formattype + " | Location: " + location + " | ISBN: " + isbn)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {

	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/print", Print)
	// http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)

}
