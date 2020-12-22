// MySQL benchmarking tool
package main

import (
	"fmt"
	"os"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

/*
	the startup script
*/
func main() {
	total := 0
	for {
		total += 1
		fmt.Println(total)
		db, err := sql.Open("mysql", os.Getenv("SQL_CONN_STRING"))
		_, err2 := db.Query("SELECT * FROM test;")
		if err != nil {
			fmt.Println(err)
			return
		}
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
}