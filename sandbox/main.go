package main

import (
	"fmt"
	"log"

	"github.com/mrtyunjaygr8/http_server_golang/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database ensured!")
	user, err := c.CreateUser("test", "test", "test", 12)
	fmt.Println(user)
}
