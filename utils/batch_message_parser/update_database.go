package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/mattn/go-sqlite3"
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

	for _, r := range recs {
		id, text, name := r[0], r[1], r[2]
		stmt := fmt.Sprintf("INSERT INTO Messages (userid, text, username) VALUES (%s, %s, %s)", id, text, name)
		log.Println(stmt)
		_, err := db.Exec(stmt)
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
