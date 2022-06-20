package main

import (
	"log"

	"github.com/mrtyunjaygr8/http_server_golang/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("database ensured!")
	user, err := c.CreateUser("test2", "test", "test", 12)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(user)
	}

	newUser, err := c.UpdateUser("test2", "test1", "test1", 121)
	log.Println(newUser)

	getUser, err := c.GetUser("test2")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(getUser)
	}
	getNotUser, err := c.GetUser("test21")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(getNotUser)
	}

	err = c.DeleteUser("test2")
	if err != nil {
		log.Println(err)
	}
	err = c.DeleteUser("test21")
	if err != nil {
		log.Println(err)
	}
}
