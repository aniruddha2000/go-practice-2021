package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(&store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
