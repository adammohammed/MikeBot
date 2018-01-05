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
	msgFile := "/home/adam/go/src/github.com/adammohammed/MikeBot/messages.csv"
	dbFile := "/home/adam/go/src/github.com/adammohammed/MikeBot/messages.db"
	f, err := os.Open(msgFile)
	errorFatal(err, "Couldn't open message file")
	defer f.Close()

	db, err := sql.Open("sqlite3", dbFile)
	errorFatal(err, "Couldn't open database")
	defer db.Close()

	reader := csv.NewReader(f)
	recs, err := reader.ReadAll()
	errorFatal(err, "Error reading csv file")

	stmt, err := db.Prepare("INSERT INTO Messages (userid, text, username) VALUES (?, ?, ?)")
	defer stmt.Close()
	errorFatal(err, "Couldn't prepare statement")
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

	err = os.Remove(msgFile)
	errorFatal(err, "Couldn't remove the message file")
	/*
		str, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalln("Error Reading from file")
		}

		contents := string(str)

		log.Println(contents)
	*/
}

func errorFatal(err error, msg string) {
	if err != nil {
		log.Fatalln(msg)
	}
}
