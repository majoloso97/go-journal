package main

import (
	"majoloso97/go-journal/cli"
	"majoloso97/go-journal/db"
)

func main() {
	connection := db.GetDBConnection()
	db.RunMigrations(connection)
	cli.HandleCLIOptions(connection)
}
