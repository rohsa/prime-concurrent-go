package main

import (
	"github.com/rohsa/prime-concurrent-go/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.ViewHandler)
	http.HandleFunc("/prime/", handlers.PrimeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
