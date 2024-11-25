package main

import (
	"os"
	"typehero_server/database"
	"typehero_server/server"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "db.db"
	}
	db, err := database.InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}
	err = server.StartServer(db)
	if err != nil {
		panic(err)
	}
}
