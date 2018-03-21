package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"
	"github.com/adammohammed/groupmebot"
	_ "github.com/mattn/go-sqlite3"
	"regexp"
	"strings"
)

/*
 Test hook functions
 Each hook should match a certain string, and if it matches
 it should return a string of text
 Hooks will be traversed until match occurs
*/
func hello(msg groupmebot.InboundMessage) string {
	resp := fmt.Sprintf("Hi, %v.", msg.Name)
	return resp
}

func hello2(msg groupmebot.InboundMessage) string {
	resp := fmt.Sprintf("Hello, %v.", msg.Name)
	return resp
}

func nameism(msg groupmebot.InboundMessage) string {
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		fmt.Println("FAILED TO GET DB")
		return ""
	}
	defer db.Close()
	re := regexp.MustCompile("(?P<name>[a-zA-Z]+)ism")
	match := re.FindStringSubmatch(msg.Text)

	if len(match) > 0 {
		name := strings.ToLower(match[1])
		fmt.Printf("Looking for message from %s\n", name)
		query := "SELECT Messages.text FROM Messages JOIN Users ON Messages.userid = Users.userid WHERE Users.name IS ? ORDER BY RANDOM() LIMIT 1;"
		var randomMessage string
		err = db.QueryRow(query, name).Scan(&randomMessage)
		switch {
		case err == sql.ErrNoRows:
			return ""
		case err != nil:
			return ""
		default:
			return randomMessage
		}

	}
	fmt.Println("FAILED TO GET MSG")
	return ""

}

func main() {

	bot, err := groupmebot.NewBotFromJson("mybot_cfg.json")
	if err != nil {
		log.Fatal("Could not create bot structure")
	}

	// Make a list of functions
	bot.AddHook("Hi!$", hello)
	bot.AddHook("Hello!$", hello2)
	bot.AddHook("[a-zA-Z]+ism", nameism)

	// Create Server to listen for incoming POST from GroupMe
	log.Printf("Listening on %v...\n", bot.Server)
	http.HandleFunc("/", bot.Handler())
	log.Fatal(http.ListenAndServe(bot.Server, nil))
}
