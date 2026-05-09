package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 3031, "Port to serve")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /")
		_, _ = fmt.Fprintf(w, "%d Hello World!\n", port)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("Starting server at %s...\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
