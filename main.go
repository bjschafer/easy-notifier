package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := NewRouter()

	port := flag.Int("port", 8080, "Port to run on")
	flag.Parse()

	if *port < 1 || *port > 65536 {
		log.Fatal(string(*port) + " is not a valid port number. Exiting.")
	}

	portString := ":" + strconv.Itoa(*port)

	fmt.Printf("Starting server on port %s", portString)

	log.Fatal(http.ListenAndServe(portString, router))
}
