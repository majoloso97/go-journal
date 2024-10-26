package cli

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"majoloso97/go-journal/db"
	"os"
)

func HandleCLIOptions(connection *sql.DB) {
	var (
		retrieve_all bool
		create_new   bool
	)
	flag.BoolVar(&retrieve_all, "retrieve-all", false, "Retrieves all journal entries")
	flag.BoolVar(&create_new, "create-new", false, "Creates new journal entry")
	flag.Parse()
	if retrieve_all {
		retrieveAll(connection)
	}
	if create_new {
		createNewEntry(connection)
	}
}

func retrieveAll(connection *sql.DB) {
	entries, err := db.GetJournalEntries(connection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(entries)

}

func createNewEntry(connection *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write your journal entry: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v", err)
		return
	}
	input = input[:len(input)-1]
	entry, err := db.SaveEntry(connection, input)
	fmt.Printf("Your new entry is: %v", entry)

}
