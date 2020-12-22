// MySQL benchmarking tool
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/*
	the connection script
*/
func ConnTest(ConnURL string) int{
	total := 0
	for {
		total += 1
		db, err := sql.Open("mysql", ConnURL)
		_, err2 := db.Query("SELECT * FROM test;")
		if err != nil {
			fmt.Println(err)
			return total
		}
		if err2 != nil {
			fmt.Println(err2)
			return total
		}
	}
}

/*
	HTTP handler
*/
func RequestCounter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	SqlString := r.FormValue("sqlconnstring")
	fmt.Fprintf(w, strconv.Itoa(ConnTest(SqlString)))
}

/*
	start the webapp
*/
func main() {
	http.HandleFunc("/", RequestCounter)
	http.ListenAndServe(":6443", nil)	
}