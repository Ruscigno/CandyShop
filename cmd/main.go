package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ruscigno/CandyShop/pkg/http/rest"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	router := rest.InitHandler()
	log.Fatal(http.ListenAndServe(":8080", router))
}
