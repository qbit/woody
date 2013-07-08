package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"log"
	"net/http"
	"os"
)

func errr(e error, msg string) {
	if e != nil {
		log.Printf("[!]: %s - %s", msg, e)
	}
}

func ConnectDB() (*sql.DB, error) {
	conn_string := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		os.Getenv("PGDBNAME"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGSSLMODE"))

	db, err := sql.Open("postgres", conn_string)
	errr(err, "dbase couldn't open database connection")

	return db, err
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

func main() {
	db, err := ConnectDB()
	errr(err, "Can't connect")
	if err != nil {
		log.Printf("YAY")
	}

	rows, err := db.Query("select * from geography_columns")
	errr(err, "Can't query")

	fmt.Printf("%#v\n", rows)
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":8080", nil)
}
