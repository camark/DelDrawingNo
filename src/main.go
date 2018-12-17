package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DbWorker struct {
	//mysql data source name
	Dsn string
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage:del <drawingno>")
		return
	}

	dno := os.Args[1]

	dbw := DbWorker{
		Dsn: "root:123@abc@tcp(10.10.10.101:3306)/product_ent",
	}
	db, err := sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()

	var icount int

	err = db.QueryRow("select count(*) from product_jiaotu where dan_no= ?", dno).Scan(&icount)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(icount)

	if icount == 0 {
		fmt.Println("Drawing no not exist!")
		return
	}

	stmt, err := db.Prepare(`delete from product_no where dan_no=?`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(dno)

	fmt.Println("Delete Drawing No " + dno + " Success")
}
