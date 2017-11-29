package database

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "dev:qrwgDGCGnMPGHtKqQhHnZYmhXfFxHyaJsQuAxnGEv@tcp(webdbinstance.cc1j8x3t3gto.ap-northeast-2.rds.amazonaws.com:3306)/webdb")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close();
}
