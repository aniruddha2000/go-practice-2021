package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aniruddha2000/goMongoAPI/router"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at 4000...")
}
