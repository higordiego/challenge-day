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

func dbConn() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "testdb"
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", dbUser, dbPass, dbName))
	if err != nil {
		return nil, err
	}
	// Testa a conexão com o banco de dados
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		log.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	var query string
	if r.Method == "POST" {
		search := r.FormValue("search")
		// Vulnerável a SQL Injection - lembrando que isso é proposital para um CTF
		query = "SELECT id, name, email FROM users WHERE name LIKE '%" + search + "%' OR email LIKE '%" + search + "%'"
	} else {
		query = "SELECT id, name, email FROM users"
	}

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Erro ao consultar o banco de dados", http.StatusInternalServerError)
		log.Println("Erro ao executar consulta:", err)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Erro ao processar os dados", http.StatusInternalServerError)
			log.Println("Erro ao escanear linha:", err)
			return
		}
		users = append(users, user)
	}

	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		log.Println("Erro ao executar template:", err)
	}
}

func main() {
	http.HandleFunc("/", Index)
	log.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
