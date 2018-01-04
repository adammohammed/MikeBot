package main

import (
	"database/sql"
	"encoding/csv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../../messages.csv")
	if err != nil {
		log.Fatalln("Error opening file")
	}
	defer f.Close()

	db, err := sql.Open("sqlite3", "../../messages.db")
	if err != nil {
		log.Fatalln("Couldn't open database")
	}
	defer db.Close()

	reader := csv.NewReader(f)
	recs, err := reader.ReadAll()

	if err != nil {
		log.Fatalln("Error parsing records")
	}

	stmt, err := db.Prepare("INSERT INTO Dummy (userid, text, username) VALUES (?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatalln("Couldn't prepare statement")
	}
	for _, r := range recs {
		log.Println(stmt)
		_, err := stmt.Exec(stmt, r[0], r[1], r[2])
		if err != nil {
			log.Fatalln("Couldn't insert message to db")
		}
	}
	/*
		str, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalln("Error Reading from file")
		}

		contents := string(str)

		log.Println(contents)
	*/
}
