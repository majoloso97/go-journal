package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"majoloso97/go-journal/db"
	"os"
)

func main() {
	var (
		retrieve_all bool
		create_new   bool
	)
	flag.BoolVar(&retrieve_all, "retrieve-all", false, "Retrieves all journal entries")
	flag.BoolVar(&create_new, "create-new", false, "Creates new journal entry")
	flag.Parse()
	connection := db.GetDBConnection()
	if retrieve_all {
		entries, err := db.GetJournalEntries(connection)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(entries)
	}
	if create_new {
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
}
