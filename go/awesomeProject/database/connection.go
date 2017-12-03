package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

const driverName string = "mysql"

func main() {
	db, err := sql.Open(driverName, dataSourceName("textDirectory/dbconection.txt"))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT USER_NAME FROM USER_INFO WHERE USER_ID = 'kimhyunjin'").Scan(&name)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(name)

}

func dataSourceName(filename string) string  {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
