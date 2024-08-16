package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

type User struct {
	ID    int
	Name  string
	Email string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "testdb"
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", dbUser, dbPass, dbName))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	var query string
	if r.Method == "POST" {
		search := r.FormValue("search")
		// Vulnerável a SQL Injection
		query = "SELECT id, name, email FROM users WHERE name LIKE '%" + search + "%' OR email LIKE '%" + search + "%'"
	} else {
		query = "SELECT id, name, email FROM users"
	}

	rows, err := db.Query(query)
	if err != nil {
		tmpl.Execute(w, nil)
	}

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	tmpl.Execute(w, users)
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
