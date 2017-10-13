package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type (
	server struct{}
)

const (
	port = "8000"
)

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "hostname: %s", hostname)
}

func main() {
	h := &server{}

	log.Fatal(http.ListenAndServe(":"+port, h))
}
