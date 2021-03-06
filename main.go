package main

import (
	"log"
	"net/http"

	"github.com/roshniashok/crud/routers"
)

func main() {

	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
