package main

import (
	"database/sql"
	"fmt"
	"time"

	// "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func main(){
	// sql.Open()
	time.NewTicker()
	db,err := sql.Open("mysql","root:123qwe@tcp(localhost:3306)/dispute_sim")
	if err!=nil{
		log.Fatal(err,"line 17")
	}
	defer db.Close()
	rows,err := db.Query("select count(*) from disputecase")
	if err!=nil{
		log.Fatal(err,"line 22")
	}
	defer rows.Close()
	var size int
	for rows.Next(){
		err := rows.Scan(&size)
		if err!=nil{
			log.Fatal(err,"line 29")
		}
		fmt.Println(size)
	}
}