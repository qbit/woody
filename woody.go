package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func errr(e error, msg string) {
	if e != nil {
		log.Printf("[!]: %s - %s", msg, e)
	}
}


func getFile(file string) []byte {
	filename := "public/" + file
	log.Printf("%s", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		data = []byte("404 - Not Found!")
	}
	errr(err, "Can't read file!")

	return data
}

func ConnectDB() (*sql.DB, error) {
	conn_string := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		os.Getenv("PGDBNAME"),
		orit(os.Getenv("PGUSER"), "postgres"),
		os.Getenv("PGPASSWORD"),
		orit(os.Getenv("PGHOST"), "localhost"),
		orit(os.Getenv("PGPORT"), "5432"),
		orit(os.Getenv("PGSSLMODE"), "disable"))

	db, err := sql.Open("postgres", conn_string)
	errr(err, "dbase couldn't open database connection")

	return db, err
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", getFile(r.URL.Path[1:]))
}

func main() {
	db, err := ConnectDB()
	errr(err, "Can't connect")
	if err != nil {
		log.Printf("YAY")
	}

	rows, err := db.Query("select * from hold_types")
	errr(err, "Can't query")

	fmt.Printf("%#v\n", rows)
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":8080", nil)
}
