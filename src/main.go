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

	if icount == 0 {
		fmt.Println("Drawing no not exist!")
		return
	}

	//fmt.Println(icount)
	var drawGuid string
	err = db.QueryRow("select guid from product_jiaotu where dan_no= ?", dno).Scan(&drawGuid)

	if err != nil {
		log.Fatal(err)
		return
	}



	stmt, err := db.Prepare(`delete from product_no where dan_no=?`)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(dno)

	stmt_delguid, err := db.Prepare(`delete from product_tuzhi where sguid=?`)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer stmt_delguid.Close()

	_, err = stmt.Exec(drawGuid)

	fmt.Println("Delete Drawing No " + dno + " Success")
}
