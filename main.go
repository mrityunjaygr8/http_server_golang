package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mrtyunjaygr8/http_server_golang/internal/database"
	"github.com/mrtyunjaygr8/http_server_golang/utils"
)

func main() {
	fmt.Println("yo")
	const addr = "localhost:8080"
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", testHandler)
	serveMux.HandleFunc("/err", testErrorHandler)

	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func testErrorHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 404, fmt.Errorf("not found"))
}
