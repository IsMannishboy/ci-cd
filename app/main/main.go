package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"

	_ "github.com/lib/pq" // важно: импорт драйвера
)

func MainHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`select name from list`)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		var items []string
		for rows.Next() {
			var item string
			err = rows.Scan(&item)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			items = append(items, item)
		}
		tmpl, err := template.ParseFiles("./html/main.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, items)

	}
}
func AddNewItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(string(body))
		_, err = db.Exec(`insert into list (namee) values ($1)`, string(body))
		if err != nil {
			fmt.Println("err:", err)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
func DeleteItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("err:", err.Error())
			return
		}
		fmt.Println(body)
		_, err = db.Exec(`delete from list where namee=$1`, string(body))
		if err != nil {
			fmt.Println(err)
		}
	}
}
func main() {
	connStr := "postgres://21savgae:1234@localhost:5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("starting server")
	http.HandleFunc("/main", MainHandler(db))
	http.HandleFunc("/add", AddNewItem(db))
	http.HandleFunc("/delete", DeleteItem(db))
}
