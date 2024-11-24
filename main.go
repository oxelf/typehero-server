package main

import (
	"fmt"
	"typehero_server/database"
	"typehero_server/server"
)

func main() {
    name := "linus"
    db, err := database.InitDatabase()
    if err != nil {
        panic(err)
    }
    fmt.Printf("hello world %s \n", name)
    server.StartServer(db)
}
