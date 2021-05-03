package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Hello, WWW!</H1>")
}

func main() {
	// registers the handler function for the given pattern in the DefaultServeMux
	http.HandleFunc("/", rootHandler)

	// listens on the TCP network address addr to handle requests on incoming connections.
	log.Fatal(http.ListenAndServe(":8080", nil)) // HL
}
