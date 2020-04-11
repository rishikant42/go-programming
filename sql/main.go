package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("vim-go")
	// connStr := "user=postgres dbname=message_config password=password host=localhost port=5432 sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// abc := "hello"

	rows, err := db.Query("select * from whatsappsandboxtemplate")
	fmt.Println(rows)
}
