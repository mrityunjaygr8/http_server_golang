package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("yo")
	const addr = "localhost:8080"
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", testHandler)

	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
