package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
}

type HogeHandler struct{}
type FugaHandler struct{}

func (h *HogeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"root",
		"",
		"localhost",
		"3306",
		"test",
	)
	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	rows, err := connection.Query("select * from todos;")

	todos := []Todo{}

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		todo := Todo{}
		rows.Scan(&todo.ID, &todo.Title)
		todos = append(todos, todo)
	}

	rows.Close()

	response, _ := json.Marshal(todos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *FugaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fuga")
}

func main() {
	hoge := HogeHandler{}
	fuga := FugaHandler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: nil, // DefaultServeMux を使用
	}

	// DefaultServeMux にハンドラを付与
	http.Handle("/hoge", &hoge)
	http.Handle("/fuga", &fuga)

	server.ListenAndServe()
}