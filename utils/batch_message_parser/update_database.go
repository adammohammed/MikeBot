package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
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

	stmt, err := db.Prepare("INSERT INTO Messages (userid, text, username) VALUES (?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatalln("Couldn't prepare statement")
	}
	for _, r := range recs {
		userid := string(r[0])
		text := string(r[1])
		name := string(r[2])
		textstmnt := fmt.Sprintf("INSERT INTO Messages (userid, text, usename) VALUES (%s, %s, %s)", userid, text, name)
		log.Println(textstmnt)
		_, err := stmt.Exec(userid, text, name)
		if err != nil {
			log.Fatalln("Couldn't insert message to db")
		}
	}
	err := os.Remove("../../messages.csv")
	if err != nil {
		Log.Fatalln("Couldn't remove messages.csv")
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
