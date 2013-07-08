package main

import (
	"database/sql"
	"net/http"
	"os"
	"log"
	"fmt"
	_ "github.com/bmizerany/pq"
)

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
	if err != nil {
		log.Printf("[!] dbase couldn't open database connection: %s", err)
		db = nil
	}

	return db, err
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Printf("YAY")
	}

	rows, err := db.Query("select * from geography_columns")
	if err != nil {
		log.Printf("%s", err)
	}

	fmt.Printf("%#v\n", rows)
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":8080", nil);
}
